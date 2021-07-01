package memory

import (
	"fmt"
	"sync"

	"github.com/seggga/cropurl/internal/storage/model"
)

// MemStorage ...
type MemStorage struct {
	DataMap map[string]*redirectData
	mu      sync.RWMutex
}

type redirectData struct {
	longURL     string
	description string
	count       int
}

func New(configPath string) (*MemStorage, error) {

	// there is no need to use a config for this storage type

	memStor := new(MemStorage)
	memStor.DataMap = map[string]*redirectData{
		"asdf": &redirectData{
			longURL:     "http://google.com",
			description: "let's pretend you have forgotten google's web-address",
			count:       0,
		},
		"qwerty": &redirectData{
			longURL:     "http://yandex.ru",
			description: "well, sometimes it's not easy to type 'yandex.ru'",
			count:       0,
		},
	}

	return memStor, nil
}

// Delete method ...
func (ms *MemStorage) Delete(shortID string) error {
	if ms.IsSet(shortID) {
		ms.mu.Lock()
		defer ms.mu.Unlock()

		delete(ms.DataMap, shortID)
		return nil
	}
	return fmt.Errorf("given short URL was not found (%s)", shortID)
}

func (ms *MemStorage) Close() {
	// well, there is nothing to do to close map
}

// IsSet checks if requested data is located in the storage
func (ms *MemStorage) IsSet(shortURL string) bool {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	_, ok := ms.DataMap[shortURL]
	return ok
}

// AddURL method adds a new redirect data (short -> long URL)
func (ms *MemStorage) AddURL(ld *model.LinkData) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.DataMap[ld.ShortID] = &redirectData{
		longURL:     ld.LongURL,
		description: ld.Description,
		count:       0,
	}
	return nil
}

// Resolve method finds long URL that correstponds to given short ID
func (ms *MemStorage) Resolve(ShortID string) (string, error) {
	// find data in the storage
	ms.mu.RLock()
	foundData, ok := ms.DataMap[ShortID]
	ms.mu.RUnlock()

	// data was not found
	if !ok {
		return "", fmt.Errorf("there is no long URL on given short URL (%s)", ShortID)
	}

	// produce increment on counter
	ms.mu.Lock()
	ms.DataMap[ShortID] = &redirectData{
		longURL:     foundData.longURL,
		description: foundData.description,
		count:       foundData.count + 1,
	}
	ms.mu.Unlock()

	return foundData.longURL, nil
}

// ViewStat method returns data from the database about given short ID
func (ms *MemStorage) ViewStat(ShortID string) (*model.LinkData, error) {
	// find data in the storage
	ms.mu.RLock()
	foundData, ok := ms.DataMap[ShortID]
	ms.mu.RUnlock()

	// data was not found
	if !ok {
		return nil, fmt.Errorf("there is no data about given short URL %s", ShortID)
	}

	return &model.LinkData{
		LongURL:     foundData.longURL,
		ShortID:     ShortID,
		Statistics:  int64(foundData.count),
		Description: foundData.description,
	}, nil
}

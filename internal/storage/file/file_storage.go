package file

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// FileStorage ...
type FStorage struct {
	FileName    string `toml:"FILE_NAME"`
	FileHandler os.File
}

func New(configPath string) (*FStorage, error) {
	fileStore := new(FStorage)
	_, err := toml.DecodeFile(configPath, fileStore)
	if err != nil {
		return nil, fmt.Errorf("error reading TOML config for file-storage, %w", err)
	}

	return fileStore, nil
}

func (fst *FStorage) Login() error {
	return nil
}

func (fst *FStorage) Logout() error {
	return nil
}

func (fst *FStorage) NewShort() error {
	return nil
}

func (fst *FStorage) Resolve() error {
	return nil
}

func (fst *FStorage) Delete() error {
	return nil
}

func (fst *FStorage) ViewInfo() error {
	return nil
}

func (fst *FStorage) Close() {
}

package file

import (
	"os"
)

// FileStorage ...
type FStorage struct {
	FileName    string `toml:"FILE_NAME"`
	FileHandler os.File
}

func New(filePath string) (*FStorage, error) {
	fileStore := new(FStorage)
	fileStore.FileName = filePath

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

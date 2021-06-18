package file

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// FileStorage ...
type FileStorage struct {
	FileName    string `toml:"FILE_NAME"`
	FileHandler os.File
}

func New(configPath string) (*FileStorage, error) {

	fileStore := new(FileStorage)
	_, err := toml.DecodeFile(configPath, fileStore)
	if err != nil {
		return nil, fmt.Errorf("error reading TOML config for file-storage, %w", err)
	}

	return fileStore, nil
}

func (fst *FileStorage) Login() error {
	return nil
}

func (fst *FileStorage) Logout() error {
	return nil
}

func (fst *FileStorage) NewShort() error {
	return nil
}

func (fst *FileStorage) Resolve() error {
	return nil
}

func (fst *FileStorage) Delete() error {
	return nil
}

func (fst *FileStorage) ViewInfo() error {
	return nil
}

func (fst *FileStorage) Close() {
}

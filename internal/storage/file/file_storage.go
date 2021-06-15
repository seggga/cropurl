package file

import "os"

// FileStorage ...
type FileStorage struct {
	FileName    string
	FileHandler os.File
}

func New() (*FileStorage, error) {
	
	return &FileStorage{
		FileName: ,
	}, nil
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

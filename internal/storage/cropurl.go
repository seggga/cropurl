package storage

import "github.com/seggga/cropurl/internal/storage/model"

// CropURLStorage describes storage neccessary methods to work with the application
type CropURLStorage interface {
	Close()

	IsSet(string) bool
	AddURI(*model.LinkData) error
	Resolve(string) (string, error)
}

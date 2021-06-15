package storage

// CropURLStorage describes storage neccessary methods to work with the application
type CropURLStorage interface {
	Close()
	Login() error
	Logout() error
	NewShort() error
	Resolve() error
	Delete() error
	ViewInfo() error
}

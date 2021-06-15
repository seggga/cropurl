package postgres

// PgStorage ...
type PgStorage struct {
}

func New() (*PgStorage, error) {
	return &PgStorage{}, nil
}

func (pgs *PgStorage) Login() error {
	return nil
}

func (pgs *PgStorage) Logout() error {
	return nil
}

func (pgs *PgStorage) NewShort() error {
	return nil
}

func (pgs *PgStorage) Resolve() error {
	return nil
}

func (pgs *PgStorage) Delete() error {
	return nil
}

func (pgs *PgStorage) ViewInfo() error {
	return nil
}

func (pgs *PgStorage) Close() {
}

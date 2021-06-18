package postgres

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/seggga/cropurl/internal/storage/model"
)

// PgStorage ...
type PgStorage struct {
	conn string `toml:"POSTGRES_CONN"`
}

func New(configPath string) (*PgStorage, error) {
	pgStor := new(PgStorage)
	_, err := toml.DecodeFile(configPath, pgStor)
	if err != nil {
		fmt.Printf("%v\n%v\n", err, pgStor.conn)
		return nil, fmt.Errorf("error reading TOML config for Postgres storage, %w", err)
	}
	return pgStor, nil
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

func (pgs *PgStorage) Delete() error {
	return nil
}

func (pgs *PgStorage) ViewInfo() error {
	return nil
}

func (pgs *PgStorage) Close() {
}

func (pgs *PgStorage) IsSet(str string) bool {
	return false
}
func (pgs *PgStorage) AddURI(ld *model.LinkData) error {
	return nil
}
func (pgs *PgStorage) Resolve(str string) (string, error) {
	return "123", nil
}

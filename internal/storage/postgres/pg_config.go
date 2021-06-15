package postgres

type PgConfig struct {
	BindAddr string `toml:"PGCONN"`
}

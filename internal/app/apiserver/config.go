package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig creates Config structure
func NewConfig() *Config {

	return &Config{
		BindAddr: "*:8080",
	}
}

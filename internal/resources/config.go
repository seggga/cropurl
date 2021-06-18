package resources

import (
	"database/sql"
	"fmt"

	"github.com/seggga/cropurl/internal/storage"

	"github.com/kelseyhightower/envconfig"
	//"go.uber.org/zap"
)

type Resources struct {
	Config Config
	Conn   *sql.DB
}

type Config struct {
	//	DiagPort    int `envconfig:"DIAG_PORT" default:"8081" required:"true"`
	// LogLevel    string `envconfig:"LOG_LEVEL" default:"info" required:"false"`
	RESTAPIPort int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL       string `envconfig:"DB_URL" default:"postgres://user:password@localhost:5432/petstore?sslmode=disable" required:"true"`
}

func New(storage storage.CropURLStorage) (*Resources, error) {
	// parse flags
	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	// create connection to DB
	conn, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		return nil, err
	}
	// check DB alive
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return &Resources{
		Config: conf,
		Conn:   conn,
	}, nil
}

// close opened resources
func (r *Resources) Release() error {
	// close DB connection
	r.Conn.Close()
	return nil
}

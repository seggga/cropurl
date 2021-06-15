package restapi

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	//"github.com/gorilla/mux"
	"github.com/go-chi/chi/v5"
	"github.com/BurntSushi/toml"

	"go.uber.org/zap"
)

// RESTAPI represents a REST API business logic server
type RESTAPI struct {
	server http.Server
	errors chan error
	logger *zap.SugaredLogger
}

type APIServer http.Server {
	Addr string `toml:"API_ADDR"`
}

// New returns a new instance of the REST API server
func New(logger *zap.SugaredLogger) (*RESTAPI, error) {

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	var apiServer APIServer
	a, err := toml.Decode("config.toml", &apiServer)
	if err != nil {
		// handle error
		return nil, err
	}

	return &RESTAPI{
		server: http.Server{
			Handler: router,
		},
		errors: make(chan error, 1),
		logger: logger,
	}, nil
}

func (rapi *RESTAPI) Start() {
	go func() {
		rapi.errors <- rapi.server.ListenAndServe()
		close(rapi.errors)
	}()
}

// Stop server.
func (rapi *RESTAPI) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rapi.server.Shutdown(ctx)
}

// Notify returns a channel to notify the caller about errors.
// If you receive an error from the channel you should stop the application.
func (rapi *RESTAPI) Notify() <-chan error {
	return rapi.errors
}

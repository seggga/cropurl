package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/seggga/cropurl/internal/restapi"
	stor "github.com/seggga/cropurl/internal/storage/memory" // storage = map in memory

	"go.uber.org/zap"
)

var (
	srvAddr string
)

func init() {
	flag.StringVar(&srvAddr, "address", "localhost:8080", "API server address and port number")
	//	flag.StringVar(&pgAddr, "postgres", "localhost:12345", "address and port for Postgres connection")
}

func main() {
	flag.Parse()
	// logger init
	logger, _ := zap.NewProduction() // nolint:errcheck : errors are unlikely while working with STDOUT
	defer func() {
		_ = logger.Sync()
	}()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")

	// initializing storage
	slogger.Info("Configuring and initializing storage...")
	storage, err := stor.New()
	if err != nil {
		slogger.Errorw("error creating storage", err)
	}

	// initializing API server
	slogger.Info("Configuring REST API server...")
	rapi, err := restapi.New(slogger, storage, srvAddr)
	if err != nil {
		slogger.Errorw("error configuring REST API server", err)
	}

	slogger.Info("Starting REST API server...")

	rapi.Start() // nolint:errcheck : errors are hendled with the channel and will cause stopping the application
	slogger.Info("The application is ready to serve requests.")

	// waiting for events to stop API server
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case x := <-interrupt:
		slogger.Infow("Received a signal.", "signal", x.String())
	case err := <-rapi.Notify():
		slogger.Errorw("Received an error from the business logic server.", "err", err)
	}

	slogger.Info("Stopping the servers...")
	err = rapi.Stop()
	if err != nil {
		slogger.Error("Got an error while stopping the business logic server.", "err", err)
	}

	slogger.Info("The app is calling the last defers and will be stopped.")
}

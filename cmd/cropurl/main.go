package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/seggga/cropurl/internal/resources"
	"github.com/seggga/cropurl/internal/restapi"
	stor "github.com/seggga/cropurl/internal/storage/postgres" // storage = postgres

	"go.uber.org/zap"
)

func main() {

	// logger init

	//nolint:errcheck Не может быть ошибки, т.к. работаем с stdout
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")
	slogger.Info("Reading configuration and initializing resources...")

	storage, err := stor.New()

	rsc, err := resources.New(slogger, storage)
	if err != nil {
		slogger.Panic("Can't initialize resources.", "err", err)
	}

	defer func() {
		err = rsc.Release()
		if err != nil {
			slogger.Errorw("Got an error during resources release.", "err", err)
		}
	}()

	slogger.Info("Configuring the application units...")
	slogger.Info("Starting the server...")
	rapi := restapi.New(slogger)
	rapi.Start()
	slogger.Info("The application is ready to serve requests.")

	// stopping API server
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case x := <-interrupt:
		slogger.Infow("Received a signal.", "signal", x.String())
	// case err := <-diag.Notify():
	// 	slogger.Errorw("Received an error from the diagnostics server.", "err", err)
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

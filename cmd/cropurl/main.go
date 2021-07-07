package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/seggga/cropurl/internal/restapi"
	stor "github.com/seggga/cropurl/internal/storage/memory" // storage = map in memory

	"go.uber.org/zap"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to configuration TOML-file")
}

func main() {
	flag.Parse()
	fmt.Printf("%v\n", configPath)
	// logger init
	//nolint:errcheck : errors are unlikely while working with STDOUT
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")

	// initializing storage
	slogger.Info("Configuring and initializing storage...")
	storage, err := stor.New(configPath)
	if err != nil {
		slogger.Errorw("error reading config for storage", err)
	}
	// rsc, err := resources.New(storage)
	// if err != nil {
	// 	slogger.Errorw("Can not initialize storage.", err)
	// }
	// defer func() {
	// 	err = rsc.Release()
	// 	if err != nil {
	// 		slogger.Errorw("error during storage release.", "err", err)
	// 	}
	// }()

	// initializing API server
	slogger.Info("Configuring REST API server...")
	rapi, err := restapi.New(slogger, storage, configPath)
	if err != nil {
		slogger.Errorw("error configuring REST API server", err)
	}

	slogger.Info("Starting REST API server...")
	//nolint:errcheck : errors are hendled with the channel and will cause stopping the application
	rapi.Start() //nolint
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

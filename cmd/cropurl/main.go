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
	"go.uber.org/zap/zapcore"
)

var (
	srvAddr  string
	logLevel string
)

func init() {
	flag.StringVar(&srvAddr, "address", ":8080", "API server address and port number")
	flag.StringVar(&logLevel, "loglevel", "info", "logging level (possible values are info, error, debug")
	//	flag.StringVar(&pgAddr, "postgres", "localhost:12345", "address and port for Postgres connection")
}

func main() {
	flag.Parse()
	addr := os.Getenv("ADDR")
	if addr != "" {
		srvAddr = addr
	}

	// create zap logger
	cfg := loggerInit(logLevel)
	logger, err := cfg.Build()
	if err != nil {
		fmt.Printf("cannot initialize logger, program exit %v", err)
		return
	}
	defer func() {
		_ = logger.Sync()
	}()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")
	slogger.Infof("env server address: %s", srvAddr)
	slogger.Infof("env log level: %s", logLevel)

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

// create a zap config structure.
func loggerInit(logLevel string) *zap.Config {
	zapLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	switch logLevel {
	case "info":
		zapLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "error":
		zapLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "debug":
		zapLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	return &zap.Config{
		Encoding:    "json",
		Level:       zapLevel,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
}

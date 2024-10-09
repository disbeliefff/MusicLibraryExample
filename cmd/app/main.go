package main

import (
	"MusicLibrary/api/server"
	"MusicLibrary/cfg"
	"errors"
	"github.com/phsym/console-slog"
	"log/slog"
	"os"
)

func main() {
	log := initLogger()
	slog.SetDefault(log)

	conf, err := cfg.LoadConfig("./cfg/app.env")
	if err != nil {
		log.Error("Failed to load config", "err", err)
	}

	s := server.NewServer(conf, log)

	err = s.Start()
	if err != nil {
		log.Error("Failed to start server", "err", err)
	}
}

func initLogger() *slog.Logger {
	logger := slog.New(
		console.NewHandler(os.Stderr, &console.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}),
	)
	slog.SetDefault(logger)
	slog.Info("Hello world!", "foo", "bar")
	slog.Debug("Debug message")
	slog.Warn("Warning message")
	slog.Error("Error message", "err", errors.New("the error"))

	logger = logger.With("foo", "bar").
		WithGroup("the-group").
		With("bar", "baz")

	logger.Info("group info", "attr", "value")

	return logger
}

package main

import (
	"MusicLibrary/api/server"
	"MusicLibrary/cfg"
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
	return logger
}

package main

import (
	"github.com/shinoxzu/news-provider-api/config"
	"github.com/shinoxzu/news-provider-api/handlers"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	config, err := config.LoadConfig(configPath)
	if err != nil {
		slog.Error("cannot load config", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Debug("config loaded", slog.Any("config", config))

	mux := http.NewServeMux()

	mux.Handle("GET /news/ria", handlers.RiaHandler{Config: config})
	mux.Handle("GET /news/tass", handlers.TassHandler{Config: config})
	
	slog.Info("running server at localhost:8080")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		slog.Error("cannot run server", slog.Any("err", err))
		os.Exit(1)
	}
}

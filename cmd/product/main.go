package main

import (
	"log/slog"
	"os"
	"product/internal/config"
	"product/internal/sl"
	"product/internal/storage"

	"github.com/go-chi/chi"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	log.Info("Starting app: 'Product'", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	st, err := storage.New(log)

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	defer storage.Close(st)
	log.Debug("Successfully connected!")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)

	// todo: init server

}

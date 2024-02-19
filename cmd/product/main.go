package main

import (
	"log/slog"
	"os"
	"product/internal/config"
	"product/internal/generator"
	"product/internal/product"
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

	s, err := storage.New(log)

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	defer s.Close()
	log.Debug("Successfully connected!")
	//
	if err = storage.Prepare(s); err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	log.Debug("Prepearing DB completed")

	//
	generator.Generate(log, s)

	pr, err := product.GetProduct(7, log, s)

	if err != nil {
		os.Exit(1)
	}
	log.Debug("GetProduct", "pr.Id", pr.Id)
	for _, attr := range pr.Attrs {
		log.Debug("GetProduct", attr.Key, attr.Value)
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)

	// todo: init server

}

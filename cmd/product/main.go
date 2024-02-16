package main

import (
	"fmt"
	"log/slog"
	"os"
	"product/internal/config"
	"product/internal/storage"
)

const (
	envLocal = "local"
	envDev   = "develop"
	envProd  = "prod"
)

func main() {
	os.Setenv("CONFIG_PATH", "C:/UserData/Projects/marketplace/product/config/local.yaml")

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting app: 'Product'", slog.String("env", cfg.Env))

	db, err := storage.Connect(cfg.Storage)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT ProductId from product")

	if err != nil {
		log.Error("exception while select query", "err", err)
	}

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Error("exception 2", "err", err)
		}
		names = append(names, name)
	}

	for i := 0; i < len(names); i++ {
		log.Debug(names[i])
	}

	log.Debug("Successfully connected!")

	// todo: init router chi

	// todo: init server

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	}
	return log
}

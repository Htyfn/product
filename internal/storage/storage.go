package storage

import (
	"database/sql"
	"log/slog"
	"os"
	"product/internal/sl"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(log *slog.Logger) (*Storage, error) {

	psqlInfo := os.Getenv("STORAGE_CONFIG")
	log.Debug("REMOVE AFTER DEBUG", "psqlInfo", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Error("Failed to open connecton to db", sl.Err(err))
		return &Storage{}, err
	}

	err = db.Ping()
	if err != nil {
		log.Error("Failed to Ping to db", sl.Err(err))
		return &Storage{}, err
	}

	return &Storage{db: db}, nil
}

func Close(storage *Storage) {
	storage.db.Close()
}

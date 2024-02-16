package storage

import (
	"database/sql"
	"fmt"
	"product/internal/config"

	_ "github.com/lib/pq"
)

func Connect(storage config.Storage) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		storage.Host, storage.Port, storage.User, storage.Password, storage.DBName)
	return sql.Open("postgres", psqlInfo)
}

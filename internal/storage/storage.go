package storage

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"product/internal/sl"

	_ "github.com/lib/pq" // init pgsql driver
)

type Storage struct {
	db *sql.DB
}

func New(log *slog.Logger) (*Storage, error) {
	const op = "storage.New"

	psqlInfo := os.Getenv("STORAGE_CONFIG") // возможно стоит вынести в константу

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Error("Failed to open connecton to db", sl.Err(err))
		return nil, fmt.Errorf("%s: %f", op, err)
	}

	err = db.Ping()
	if err != nil {
		log.Error("Failed to Ping to db", sl.Err(err))
		return nil, fmt.Errorf("%s: %f", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

// returns productId
//
//	INSERT INTO product (seller, price, curr) VALUES (seller, price, curr) RETURNING id;
func (s *Storage) SaveProduct(seller int, price float64, curr int16) (int, error) {
	const op = "storage.SaveProduct"

	stmt, err := s.db.Prepare("INSERT INTO product (seller, price, curr) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var productId int
	err = stmt.QueryRow(seller, price, curr).Scan(&productId)
	if err != nil {
		return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return productId, nil
}

// тут бы обрезать value до 4000 знаков
func (s *Storage) SaveAttrStr(productId int, key string, value string) error {
	const op = "storage.SaveAttrStr"

	_, err := s.db.Exec("INSERT INTO productAttrStr (productId, key, value) VALUES ($1, $2, $3)", productId, key, value)
	if err != nil {
		return fmt.Errorf("%s: execute statement: %w", op, err)
	}
	return nil
}

// returns seller, price, curr, err
//
//	select seller, price, curr from product where id = ?;
func (s *Storage) GetProductById(id int) (int, float64, int16, error) {
	const op = "storage.GetProductById"
	var seller int
	var price float64
	var curr int16

	stmt, err := s.db.Prepare("select seller, price::money::numeric, curr from product where id = $1 ")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	err = stmt.QueryRow(id).Scan(&seller, &price, &curr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return seller, price, curr, nil
}

// TODO
// returns sql.Rows (key, value)
func (s *Storage) GetProductAttrStr(productId int) (*sql.Rows, error) {
	const op = "storage.GetProductAttrStr"

	stmt, err := s.db.Prepare("select key, value from productAttrStr where productId = $1 ")
	if err != nil {
		return nil, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	rows, err := stmt.Query(productId)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return rows, nil
}

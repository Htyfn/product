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

// инициирует таблицы
// TODO перенести в миграции
func Prepare(storage *Storage) error {
	const op = "storage.Prepare"

	stmt, err := storage.db.Prepare(`CREATE TABLE IF NOT EXISTS 
	product (
		Id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		seller      INT NOT NULL,
		price       MONEY NOT NULL,
		curr        INT NOT NULL,
		createDate  DATE,
		updateDate  DATE
	);`)
	if err != nil {
		return fmt.Errorf("%s: %f", op+"1", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %f", op+"1", err)
	}

	stmt, err = storage.db.Prepare("CREATE INDEX IF NOT EXISTS product_idx_sellerId ON product(seller);")
	if err != nil {
		return fmt.Errorf("%s: %f", op+"2", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %f", op+"2", err)
	}

	stmt, err = storage.db.Prepare(`CREATE TABLE IF NOT EXISTS 
	productAttrStr (
		ProductId   INT references product(Id),
		Key			VARCHAR(500),
		Value	    VARCHAR(4000)
	);`)
	if err != nil {
		return fmt.Errorf("%s: %f", op+"3", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %f", op+"3", err)
	}

	stmt, err = storage.db.Prepare("CREATE INDEX IF NOT EXISTS ProductAttrStr_idx_productId ON product(Id);")
	if err != nil {
		return fmt.Errorf("%s: %f", op+"4", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %f", op+"4", err)
	}

	return nil
}

func Close(storage *Storage) {
	storage.db.Close()
}

// returns productId
func SaveProduct(storage *Storage, seller int64, price float32, curr int16) (int64, error) {
	const op = "storage.SaveProduct"

	//  INSERT INTO product (seller, price, curr) VALUES (seller, price, curr) RETURNING id;

	stmt, err := storage.db.Prepare("INSERT INTO product (seller, price, curr) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var productId int64

	err = stmt.QueryRow(seller, price, curr).Scan(&productId)
	if err != nil {
		return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return productId, nil
}

func SaveAttrStr(storage *Storage, productId int64, key string, value string) error {

	//  INSERT INTO product (seller, price, curr) VALUES (seller, price, curr) RETURNING id;

	return nil
}

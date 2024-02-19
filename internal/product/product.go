package product

import (
	"fmt"
	"log/slog"
	"product/internal/sl"
	"product/internal/storage"
)

type Product struct {
	Id     int
	Seller int
	Price  float64
	Curr   int16
	Attrs  []AttrStr
}

type AttrStr struct {
	Key   string
	Value string
}

const curr_RU = 643

func (pr *Product) SaveProduct(log *slog.Logger, s storage.Storage) error {
	const op = "SaveProduct"
	var err error

	if pr.Curr == 0 {
		pr.Curr = curr_RU
	}
	pr.Id, err = s.SaveProduct(pr.Seller, pr.Price, pr.Curr)

	if err != nil {
		log.Error("failed to SaveProduct", sl.Err(err))
		return fmt.Errorf("%s: failed to Save Product: %w", op, err)
	}

	for _, attr := range pr.Attrs {
		s.SaveAttrStr(pr.Id, attr.Key, attr.Value)
		if err != nil {
			log.Error("failed to SaveAttrStr", sl.Err(err))
			return fmt.Errorf("%s: failed to SaveAttrStr: %w", op, err)
		}
	}
	return nil
}

func GetProduct(id int, log *slog.Logger, s *storage.Storage) (*Product, error) {
	const op = "GetProduct"
	var err error
	var pr = Product{Id: id}
	pr.Seller, pr.Price, pr.Curr, err = s.GetProductById(7)

	if err != nil {
		log.Error("failed to GetProduct", sl.Err(err))
		return &Product{}, fmt.Errorf("%s: failed to GetProductById: %w", op, err)
	}

	log.Debug("GetProduct", "pr.Seller", pr.Seller, "pr.Id", pr.Id)

	attrRows, err := s.GetProductAttrStr(7)

	if err != nil {
		log.Error("failed to GetProduct", sl.Err(err))
		return &Product{}, fmt.Errorf("%s: failed to GetProductAttrStr: %w", op, err)
	}

	var attrStr AttrStr
	for attrRows.Next() {

		err = attrRows.Scan(&attrStr.Key, &attrStr.Value)
		if err != nil {
			return &Product{}, fmt.Errorf("%s: failed to parse rows: %w", op, err)
		}
		pr.Attrs = append(pr.Attrs, attrStr)
	}

	return &pr, nil
}

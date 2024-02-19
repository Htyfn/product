package generator

import (
	"log/slog"
	"os"
	"product/internal/product"
	"product/internal/sl"
	"product/internal/storage"
)

func Generate(log *slog.Logger, s *storage.Storage) {
	var attrs = []product.AttrStr{{"TITLE", "Чайник заварочный, 800 мл"}, {"TYPE", "Чайник заварочный"}, {"QUANTITY", "1"}, {"IMG", "202401042123123.webp"}, {"Материал крышки", "Чугун"}}
	var p = product.Product{Seller: 1, Price: 2748, Attrs: attrs}

	err := p.SaveProduct(log, *s)
	if err != nil {
		log.Error("failed to insert rows", sl.Err(err))
		os.Exit(1)
	}
	log.Debug("Successfully inserted into product", "productId", p.Id)

	attrs = []product.AttrStr{{"TITLE", "Батарейки ААА аккумуляторные GP"}, {"TYPE", "Аккумуляторная батарейка"}, {"QUANTITY", "6"}, {"IMG", "202402102545563.webp"}, {"Емкость, мА•ч", "950"}}
	p = product.Product{Seller: 2, Price: 1108, Attrs: attrs}
	err = p.SaveProduct(log, *s)
	if err != nil {
		log.Error("failed to insert rows", sl.Err(err))
		os.Exit(1)
	}

	log.Debug("Successfully inserted into product", "productId", p.Id)
}

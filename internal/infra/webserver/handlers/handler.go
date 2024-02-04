package handlers

import (
	"database/sql"

	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
)

type handlers struct {
	Product *ProductHandler
}

func LoadHandlers(conn *sql.DB) (*handlers, error) {
	repository, err := repositories.NewRepository(conn)
	if err != nil {
		return nil, err
	}

	productHandler := NewProductHandler(repository.Product.(*repositories.ProductRepository))

	return &handlers{
		Product: productHandler,
	}, nil
}
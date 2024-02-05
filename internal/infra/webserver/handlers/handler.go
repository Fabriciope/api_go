package handlers

import (
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
)

type handlers struct {
	Product *productHandler
	User *userHandler
}

func LoadHandlers() (*handlers, error) {
	repository, err := repositories.NewRepository()
	if err != nil {
		return nil, err
	}

	productHandler := newProductHandler(repository.Product)
	userHandler := newUserHandler(repository.User)

	return &handlers{
		Product: productHandler,
		User: userHandler,
	}, nil
}
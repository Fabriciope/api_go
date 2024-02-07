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

	return &handlers{
		Product: newProductHandler(repository.Product),
		User: newUserHandler(repository.User),
	}, nil
}
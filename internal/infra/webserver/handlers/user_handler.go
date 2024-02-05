package handlers

import (
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
	"github.com/Fabriciope/my-api/internal/services"
)

type userHandler struct {
	repository repositories.RepositoryInterface
	service *services.UserService
}

func newUserHandler(repository repositories.RepositoryInterface) *userHandler {
	return &userHandler{
		repository: repository,
		service: &services.UserService{
			Repository: repository,
		},
	}
}
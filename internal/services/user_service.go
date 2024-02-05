package services

import "github.com/Fabriciope/my-api/internal/infra/database/repositories"

type UserService struct {
	Repository repositories.RepositoryInterface
}
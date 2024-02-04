package repositories

import (
	"database/sql"

	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	Create(models.ModelInterface) error
	Update(models.ModelInterface) error
	Delete(uuid.UUID) error
	FindAllWithPagination(page, limit int, sort string) ([]models.ModelInterface, error)
	FindOneWhere(string, interface{}) (models.ModelInterface, error)
	validateModel(models.ModelInterface) bool
}


type Container struct {
	User RepositoryInterface
	Product RepositoryInterface
}

func NewRepository(conn *sql.DB) (*Container, error) {
	return &Container {
		User: newUserRepository(conn),
		Product: newProductRepository(conn),
	}, nil
}
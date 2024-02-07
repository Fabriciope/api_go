package repositories

import (
	"github.com/Fabriciope/my-api/internal/infra/database"
	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

// TODO: fazer uma função makeAction(func()) para a camada de serviço personalizar as query
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

// TODO: ter duas conexões por repositório (uma para leitura e outra para escrita)
func NewRepository() (*Container, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}

	return &Container {
		User: newUserRepository(conn),
		Product: newProductRepository(conn),
	}, nil
}
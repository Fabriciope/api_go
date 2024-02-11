package services

import (
	"errors"

	"github.com/Fabriciope/my-api/internal/dto"
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

// TODO: testar
type ProductService struct {
	Repository repositories.RepositoryInterface
}

func NewProductService(repository repositories.RepositoryInterface) *ProductService {
	return &ProductService{repository}
}

func (s *ProductService) CreateProduct(dto *dto.CreateProductInput) error {
	product, err := models.NewProduct((*dto).Name, (*dto).Price)
	if err != nil {
		return err
	}

	//TODO: verificar duplicação
	if s.Repository.Create(product) != nil {
		return errors.New("Error when inserting product: " + product.Name)
	}

	return nil
}

func (s *ProductService) UpdateProduct(id string, dto *dto.UpdateProductInput) error {
	idConverted, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}

	productFound, err := s.Repository.FindOneWhere("id", idConverted)
	if err != nil || productFound == nil {
		return errors.New("product not found")
	}

	// TODO: criar uma função na camada de dto ou de modelo (updateModelFromDTO) ou (makeModelFromDTO)
	product := productFound.(*models.Product)
	if newName := dto.Name; newName != "" {
		product.Name = newName
	}
	if newPrice := dto.Price; newPrice != 0 {
		product.Price = newPrice
	}
	//

	err = s.Repository.Update(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(id string) error {
	idConverted, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}

	productFound, err := s.Repository.FindOneWhere("id", idConverted)
	if err != nil || productFound == nil {
		return errors.New("product not found")
	}

	err = s.Repository.Delete(idConverted)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) GetAllWithPagination(page, limit int, sort string) ([]models.Product, error) {
	productsFound, err := s.Repository.FindAllWithPagination(page, limit, sort)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for i := range productsFound {
		if p, ok := productsFound[i].(*models.Product); ok {
			products = append(products, *p)
			continue
		}
		return nil, errors.New("error when converting to model")
	}

	return products, nil
}

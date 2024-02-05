package services

import (
	"errors"

	"github.com/Fabriciope/my-api/internal/dto"
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

type ProductService struct {
	Repository repositories.RepositoryInterface
}

func (s *ProductService) CreateProduct(dto *dto.CreateProductInput) error {
	product, err := models.NewProduct((*dto).Name, (*dto).Price)
	if err != nil {
		return err
	}

	//TODO: verificar duplicação
	// err = s.Repository.Create(product)
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

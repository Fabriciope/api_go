package repositories

import (
	"testing"

	"github.com/Fabriciope/my-api/configs"
	"github.com/Fabriciope/my-api/internal/infra/database"
	"github.com/Fabriciope/my-api/internal/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func init() {
	configs.Cfg = &configs.Config{
		DBDriver:   "mysql",
		DBHost:     "localhost",
		DBPort:     7000,
		DBUser:     "root",
		DBName:     "api_golang",
		DBPassword: "password",
	}

	conn, _ = database.Connect()
}

func TestInsertNewProduct(t *testing.T) {
	repository, _ := NewRepository(conn)

	product, _ := models.NewProduct(gofakeit.ProductName(), int(gofakeit.Price(800, 3000)))
	err := repository.Product.Create(product)

	assert.Nil(t, err)

	_, err = repository.Product.FindOneWhere("id", product.ID)
	
	assert.Nil(t, err)
}

func TestUpdateAProduct(t *testing.T) {
	repository, _ := NewRepository(conn)
	
	product, _ := models.NewProduct(gofakeit.ProductName(), int(gofakeit.Price(800, 3000)))
	repository.Product.Create(product)

	productFound, _ := repository.Product.FindOneWhere("id", product.ID)
	p := productFound.(*models.Product)
	p.Name = "New name"
	
	err := repository.Product.Update(p)
	
	assert.Nil(t, err)
	
	productUpdated, err := repository.Product.FindOneWhere("name", p.Name)
	
	assert.NotNil(t, productUpdated)
	assert.Nil(t, err)
}

func TestDeleteAProduct(t *testing.T) {
	repository, _ := NewRepository(conn)
	
	product, _ := models.NewProduct(gofakeit.ProductName(), int(gofakeit.Price(800, 3000)))
	repository.Product.Create(product)
	
	err := repository.Product.Delete(product.ID)
	
	assert.Nil(t, err)
	
	productDeleted, err := repository.Product.FindOneWhere("id", product.ID)
	
	assert.NotNil(t, err)
	assert.Nil(t, productDeleted)
}

func TestFindAllProductsWithPagination(t *testing.T) {
	repository, _ := NewRepository(conn)

	for i := 0; i <= 10; i++ {
		product, _ := models.NewProduct(gofakeit.Name(), int(gofakeit.Price(1000, 3000)))
		repository.Product.Create(product)
	}

	page, limit := 3, 2
	productsFound, err := repository.Product.FindAllWithPagination(page, limit, "desc")
	
	assert.Nil(t, err)
	assert.Len(t, productsFound, limit)
}

func TestFindOneProduct(t *testing.T) {
	repository, _ := NewRepository(conn)

	product, _ := models.NewProduct(gofakeit.ProductName(), int(gofakeit.Price(800, 3000)))
	repository.Product.Create(product)

	productFound, err := repository.Product.FindOneWhere("id", product.ID)

	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	p := productFound.(*models.Product)

	assert.Equal(t, product.ID, p.ID)
	assert.Equal(t, product.Name, p.Name)
	assert.Equal(t, product.Price, p.Price)
	assert.Equal(t, product.CreatedAt, p.CreatedAt)
}

func TestFindOneProductWhenDataIsWrong(t *testing.T) {
	repository, _ := NewRepository(conn)

	userFound, err := repository.Product.FindOneWhere("name", "wrong-name")

	assert.Nil(t, userFound)
	assert.NotNil(t, err)
}

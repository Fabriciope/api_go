package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

type productRepository struct {
	db *sql.DB
	defaultActions defaultActions
}

func newProductRepository(conn *sql.DB) *productRepository {
	return &productRepository{
		db: conn,
		defaultActions: defaultActions{
			db: conn,
			table: "products",
		},
	}
}

func (pr *productRepository) Create(product models.ModelInterface) error {
	if !pr.validateModel(product) {
		return errInvalidModel
	}
	
	err := pr.defaultActions.Insert(product)
	if err != nil {
		return err
	}
	
	return nil
}

func (pr *productRepository) Update(product models.ModelInterface) error {
	if !pr.validateModel(product) {
		return errInvalidModel
	}

	err := pr.defaultActions.Update(product) 
	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) Delete(id uuid.UUID) error {
	return pr.defaultActions.destroy(id)
}

func (pr *productRepository) FindAllWithPagination(page, limit int, sort string) ([]models.ModelInterface, error) {
	sort = strings.ToUpper(sort)
	if sort != "ASC" && sort != "DESC" {
		sort = "ASC"
	} 

	if page <= 0 || limit <= 0 {
		return nil, errPageOrLimitAreWrong
	}

	rows, err := pr.db.Query(fmt.Sprintf(
		"SELECT * FROM products ORDER BY created_at %s LIMIT %d OFFSET %d",
		sort, limit, (page - 1) * limit,
	))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []models.ModelInterface
	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, &p)
	}

	return products, nil
}

func (pr *productRepository) FindOneWhere(attr string, value interface{}) (models.ModelInterface, error) {
	query := fmt.Sprintf("SELECT * FROM products WHERE %s = ?", attr)
	stmt, err := pr.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p models.Product
	err = stmt.QueryRow(value).Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (pr *productRepository) validateModel(model models.ModelInterface) bool {
	_, ok := model.(*models.Product)
	return ok
}
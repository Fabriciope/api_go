package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid Price")
)

// TODO: ver se tem o campo user_id
type Product struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     int       `json:"price" db:"price"`
	CreatedAt string    `json:"create_at" db:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        uuid.New(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().String()[:19],
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if (*p).Name == "" {
		return ErrNameIsRequired
	}

	if (*p).Price == 0 {
		return ErrPriceIsRequired
	}

	if (*p).Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}

func (p *Product) Get(field string) reflect.Value {
	return reflect.ValueOf(*p).FieldByName(field)
}

func (p *Product) DataForDB() map[string]interface{} {
	data := make(map[string]interface{}, 4)
	tProduct := reflect.TypeOf(*p)
	fields := reflect.VisibleFields(tProduct)
	for _, field := range fields {
		if sField, ok := tProduct.FieldByName(field.Name); ok {
			fieldNameInDatabase := string(sField.Tag.Get("db"))
			data[fieldNameInDatabase] = p.Get(field.Name).Interface()
		}
	}

	return data
}

func (u Product) ToJson() []byte {
	responseJson, err := json.Marshal(u)
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return []byte(fmt.Sprintf(`"product": %s`, string(responseJson)))
}

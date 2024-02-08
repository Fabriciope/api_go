package dto

import (
	"encoding/json"
	"fmt"

	"github.com/Fabriciope/my-api/internal/models"
	"github.com/Fabriciope/my-api/pkg"
)

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type AllProductsOutput struct {
	Page     uint                    `json:"page"`
	Limit    uint                    `json:"limit"`
	Sort     string                  `json:"sort"`
	Products []models.ModelInterface `json:"products"`
}

func (r AllProductsOutput) ToJson() []byte {
	responseJson, err := json.Marshal(r)
	if err != nil {
		pkg.LogError("Error: marshal in AllProductsOutput from ToJson()", err)
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return []byte(fmt.Sprintf(`"products": %s`, string(responseJson)))
}

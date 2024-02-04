package responses

import (
	"encoding/json"
	"fmt"

	"github.com/Fabriciope/my-api/internal/models"
)

type AllProductsResponse struct {
	Page     uint   `json:"page"`
	Limit    uint   `json:"limit"`
	Sort     string `json:"sort"`
	Products []models.ModelInterface `json:"products"`
}

func (r AllProductsResponse) ToJson() []byte {
	responseJson, err := json.Marshal(r)
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return []byte(fmt.Sprintf(`"products": %s`, string(responseJson)))
}

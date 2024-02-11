package dto

import (
	"encoding/json"

	"github.com/Fabriciope/my-api/pkg"
)

type OutputInterface interface {
	ToJson() []byte
}

type DefaultOutput struct {
	Error   bool   `json:"error" example:"false"`
	Message string `json:"message" example:"message"`
}

func (r DefaultOutput) ToJson() []byte {
	responseJson, err := json.Marshal(r)
	if err != nil {
		pkg.LogError("Error: marshal in DefaultOutput from ToJson()", err)
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return responseJson
}

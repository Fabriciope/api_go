package dto

import (
	"encoding/json"
	"strings"

	"github.com/Fabriciope/my-api/pkg"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}

func (r GetJWTOutput) ToJson() []byte {
	responseJson, err := json.Marshal(r)
	if err != nil {
		pkg.LogError("Error: marshal in GetJWTOutput from ToJson()", err)
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return []byte(strings.TrimSuffix(strings.TrimPrefix(string(responseJson), "{"), "}"))
}

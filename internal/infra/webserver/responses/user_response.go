package responses

import (
	"encoding/json"
	"strings"
)

type GetJWTResponse struct {
	AccessToken string `json:"access_token"`
}

func (r GetJWTResponse) ToJson() []byte {
	responseJson, err := json.Marshal(r)
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return []byte(strings.TrimSuffix(strings.TrimPrefix(string(responseJson), "{"), "}"))
}

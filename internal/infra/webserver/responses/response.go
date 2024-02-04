package responses

import (
	"encoding/json"
	"fmt"
	"strings"
)

type DefaultResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func ErrorToJson(message string) (responseJson []byte) {
	responseJson, err := json.Marshal(DefaultResponse{
		Error:   true,
		Message: message,
	})
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return
}

func SuccessToJson(message string) (responseJson []byte) {
	responseJson, err := json.Marshal(DefaultResponse{
		Error:   false,
		Message: message,
	})
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)// TODO: ver se esta contradição de um erro no sucesso é semantico
	}

	return
}

type ResponseInterface interface {
	ToJson() []byte
}

func SuccessWithDataToJson(message string, data ResponseInterface) []byte {
	defaultResponse, err := json.Marshal(DefaultResponse{Error: false, Message: message})
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)// TODO: ver se esta contradição de um erro no sucesso é semantico
	}

	return  []byte(strings.Replace(string(defaultResponse), "}", fmt.Sprint(", ", string(data.ToJson())), 1))
}

func ErrorWithDataToJson() {

}


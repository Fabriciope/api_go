package handlers

import (
	"encoding/json"
	"strings"

	"github.com/Fabriciope/my-api/internal/dto"
	"github.com/Fabriciope/my-api/pkg"
)

func successWithDataToJson(message string, data dto.OutputInterface) []byte {
	defaultResponse, err := json.Marshal(dto.DefaultOutput{Error: false, Message: message})
	if err != nil {
		pkg.LogError("Error: marshal from SuccessWithDataToJson()", err)
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return []byte(strings.Replace(string(defaultResponse), "}", ", " + string(data.ToJson()) + "}", 1))
}

// TODO: make
// func errorWithDataToJson(message string, data OutputInterface) []byte {

// }

func successToJson(message string) []byte {
	responseJson, err := json.Marshal(dto.DefaultOutput{
		Error:   false,
		Message: message,
	})
	if err != nil {
		pkg.LogError("Error: marshal from SuccessToJson()", err)
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return responseJson
}

func errorToJson(message string) (responseJson []byte) {
	responseJson, err := json.Marshal(dto.DefaultOutput{
		Error:   true,
		Message: message,
	})
	if err != nil {
		pkg.LogError("Error: marshal from ErrorToJson()", err)
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return
}
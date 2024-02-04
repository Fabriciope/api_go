package models

import (
	"reflect"

	// "github.com/Fabriciope/my-api/internal/infra/webserver/responses"
)

type ModelInterface interface {
	ToJson() []byte // TODO: testar esta função nos modelos
	Get(string) reflect.Value
	// SetID(uint)
	DataForDB() map[string]interface{}
	//TODO: makeModelFromDTO
}
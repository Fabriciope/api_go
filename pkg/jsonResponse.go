package pkg

import (
	"fmt"
	// "reflect"
	// "strings"
)

// TODO: testar funções abaixo

func ErrorResponseToJson(message string) []byte {
	return []byte(fmt.Sprintf(`{"error": true, "message":"%s"}`, message))
}

func SuccessResponseToJson(message string) []byte {
	return []byte(fmt.Sprintf(`{"error": false, "message":"%s"}`, message))
}

// TODO: continuar com a ideia abaixo
// func MakeResponseToJson(data map[string]interface{}) []byte {
// 	var responseJson string
// 	for field, value := range data {
// 		typeValue := reflect.TypeOf(value).Kind()
// 		fmt.Println(typeValue)
// 		switch typeValue {
// 		case reflect.Map:
// 			mapValue := MakeResponseToJson(value.(map[string]interface{}))
// 			responseJson += string(fmt.Sprintf(`"%s": %v,`, field, mapValue))
// 		case reflect.Slice, reflect.Array:
// 			var sliceValue string
// 			for _, content := range value.([]interface{}) {
// 				if fmt.Sprintf("%T", value) == "map" {
// 					mapValue := MakeResponseToJson(content.(map[string]interface{}))
// 					sliceValue += string(fmt.Sprintf(`"%s": %v,`, field, mapValue))
// 				}
// 				sliceValue += fmt.Sprintf(`"%s": %v,`, field, value)
// 			}
// 			responseJson += fmt.Sprintf(`"%s": %v,`, field, sliceValue)
// 		}

// 		responseJson += fmt.Sprintf(`"%s": %v,`, field, value)
// 		// if typeValue == reflect.Map {
// 		// 	mapValue := MakeResponseToJson(value.(map[string]interface{}))
// 		// 	responseJson += string(fmt.Sprintf(`"%s": %v,`, field, mapValue))
// 		// } else if typeValue == reflect.Slice || typeValue == reflect.Array {
// 		// 	var sliceValue string
// 		// 	for _, content := range value.([]interface{}) {
// 		// 		if fmt.Sprintf("%T", value) == "map" {
// 		// 			mapValue := MakeResponseToJson(content.(map[string]interface{}))
// 		// 			sliceValue += string(fmt.Sprintf(`"%s": %v,`, field, mapValue))
// 		// 		}
// 		// 		sliceValue += fmt.Sprintf(`"%s": %v,`, field, value)
// 		// 	}
// 		// 	responseJson += fmt.Sprintf(`"%s": %v,`, field, sliceValue)
// 		// }
// 	}

// 	return []byte(fmt.Sprintf(`{%s}`, strings.TrimSuffix(responseJson, ",")))
// }

// func resolveMapToJson(value map[string]interface{}) {

// }

// func resolveSliceToJson(value []interface{}) {

// }

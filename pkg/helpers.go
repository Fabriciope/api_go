package pkg   

import "fmt"

func LogError(message string, err error) {
	fmt.Println(message, "\nErr:", err)
}

func Log(message ...interface{}) {
	fmt.Println(message...)
}
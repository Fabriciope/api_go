package database

import (
	"testing"

	"github.com/Fabriciope/my-api/configs"
	"github.com/stretchr/testify/assert"
)


func init() {
	configs.Cfg = &configs.Config{
		DBDriver: "mysql",
		DBHost: "localhost",
		DBPort: 7000,
		DBUser: "root",
		DBName: "api_golang",
		DBPassword: "password",
	}
}

func TestConnection(t *testing.T) {
	conn, err := Connect()

	assert.Nil(t, err)
	assert.NotNil(t,conn)
}
package database

import (
	"testing"

	"github.com/Fabriciope/my-api/configs"
	"github.com/stretchr/testify/assert"
)


func init() {
	configs.Cfg.DBDriver = "mysql"
	configs.Cfg.DBHost = "localhost"
	configs.Cfg.DBPort = 7000
	configs.Cfg.DBUser = "root"
	configs.Cfg.DBName = "api_golang"
	configs.Cfg.DBPassword = "password"	
}

func TestConnection(t *testing.T) {
	conn, err := Connect()

	assert.Nil(t, err)
	assert.NotNil(t,conn)
}
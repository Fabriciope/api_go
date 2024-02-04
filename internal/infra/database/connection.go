package database

import (
	"database/sql"
	"fmt"

	"github.com/Fabriciope/my-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	config := configs.Cfg
	conn, err := sql.Open(
		config.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
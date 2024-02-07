package database

import (
	"database/sql"
	"fmt"

	"github.com/Fabriciope/my-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

// TODO: quando fazer o commit desta parte (fix sintaxe)
func Connect() (*sql.DB, error) {
	conn, err := sql.Open(
		configs.Cfg.DBDriver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			configs.Cfg.DBUser,
			configs.Cfg.DBPassword,
			configs.Cfg.DBHost,
			configs.Cfg.DBPort,
			configs.Cfg.DBName,
		),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
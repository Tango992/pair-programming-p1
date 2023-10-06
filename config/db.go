package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/pair_programming")
	if err != nil {
		return db, err
	}
	return db, nil
}

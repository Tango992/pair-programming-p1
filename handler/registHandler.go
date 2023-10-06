package handler

import (
	"database/sql"
	"errors"
	"pair-programming/entity"
	my "github.com/go-mysql/errors"
	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User, db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?,?)", user.Username, hashedPassword)
	if my.MySQLErrorCode(err) == 1062 {
		return errors.New("username tidak tersedia. silakan pilih username lain")
	}
	return err
}
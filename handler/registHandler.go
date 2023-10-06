package handler

import (
	"database/sql"
	"pair-programming/entity"
	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User, db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?,?)", user.Username, hashedPassword)
	return err
}
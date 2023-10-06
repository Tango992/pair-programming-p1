package handler

import (
	"database/sql"
	"pair-programming/entity"
	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string, db *sql.DB) (entity.User, bool, error) {
	var user entity.User
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return user, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}
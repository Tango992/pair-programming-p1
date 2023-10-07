package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"pair-programming/entity"
	my "github.com/go-mysql/errors"
)

func AddCart(db *sql.DB, user entity.User, gameId int) error {
	_, err := db.Exec("INSERT INTO user_games (user_id, game_id) VALUES (?,?)", user.Id, gameId)
	if err != nil {
		if my.MySQLErrorCode(err) == 1452 {
			return errors.New("game id tidak terdaftar")
		}
		return err
	}

	fmt.Printf("\nGame ID %v berhasil dimasukkan ke dalam cart\n", gameId)
	return nil
}
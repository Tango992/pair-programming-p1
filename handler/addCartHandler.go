package handler

import (
	"database/sql"
	"fmt"
	"pair-programming/entity"
)

func AddCart(db *sql.DB, user entity.User, gameId int) error {
	_, err := db.Exec("INSERT INTO user_games (user_id, game_id) VALUES (?,?)", user.Id, gameId)
	fmt.Printf("Game ID %v berhasil dimasukkan ke dalam cart\n", gameId)
	return err
}
package handler

import (
	"database/sql"
	"fmt"
	"pair-programming/entity"
)

func DeleteCart(db *sql.DB, user entity.User, gameId int) error {
	_, err := db.Exec("DELETE FROM user_games WHERE user_id = ? AND game_id = ?", user.Id, gameId)
	fmt.Printf("Game ID %v berhasil dihapus dari cart\n", gameId)
	return err
}

package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"pair-programming/entity"
)

func DeleteCart(db *sql.DB, user entity.User, gameId int) error {
	result, err := db.Exec("DELETE FROM user_games WHERE user_id = ? AND game_id = ?", user.Id, gameId)
	if i, _ := result.RowsAffected(); i == 0 {
		return errors.New(fmt.Sprintf("game id %v tidak ada di dalam cart anda", gameId))
	}
	fmt.Printf("Game ID %v berhasil dihapus dari cart\n", gameId)
	return err
}
package handler

import (
	"database/sql"
	"fmt"
	"pair-programming/entity"
)

func DisplayCart(db *sql.DB) error {
	rows, err := db.Query("SELECT name, description, published FROM `user_games` JOIN games ON games.id = user_games.game_id WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer rows.Close()

	var games entity.Games
	fmt.Printf("\n[ID] - [Name] - [Description] - [Published]")
	for rows.Next() {
		err := rows.Scan(&games.Id, &games.Name, &games.Description, &games.Published)
		if err != nil {
			return err
		}

		fmt.Printf("[%v] - [%v] - [%v] - [%v]\n", games.Id, games.Name, games.Description, games.Published)
	}
	fmt.Println()
	return nil
}

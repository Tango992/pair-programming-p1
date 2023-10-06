package handler

import (
	"database/sql"
	"fmt"
	"pair-programming/entity"
)

func DisplayCart(db *sql.DB, user entity.User) error {
	rows, err := db.Query("SELECT name, description, published FROM `user_games` JOIN games ON games.id = user_games.game_id WHERE user_id = ?", user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var displayCart entity.DisplayCart
	fmt.Printf("\n[Name] - [Description] - [Published]\n")
	for rows.Next() {
		err := rows.Scan(&displayCart.Name, &displayCart.Description, &displayCart.Published)
		if err != nil {
			return err
		}

		fmt.Printf("[%v] - [%v] - [%v]\n", displayCart.Name, displayCart.Description, displayCart.Published)
	}
	fmt.Println()
	return nil
}

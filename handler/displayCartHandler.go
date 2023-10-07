package handler

import (
	"database/sql"
	"fmt"
	"pair-programming/entity"
)


func DisplayCart(db *sql.DB, user entity.User) error {
	rows, err := db.Query("SELECT game_id, name, description, published FROM `user_games` JOIN games ON games.id = user_games.game_id WHERE user_id = ?", user.Id)
	if err != nil {
		return err
	}

	carts, err := GetCartFromDb(rows)
	if err != nil {
		return err
	}

	if len(carts) < 1 {
		fmt.Printf("\nCart anda kosong. Silakan gunakan fitur tambah cart\n")
		return nil
	}

	fmt.Printf("\n[ID] - [Name] - [Description] - [Published]\n")
	for _, cart := range carts {
		fmt.Printf("[%v] - [%v] - [%v] - [%v]\n", cart.Id, cart.Name, cart.Description, cart.Published)
	}
	return nil
}


func GetCartFromDb(rows *sql.Rows) ([]entity.DisplayCart, error) {
	defer rows.Close()
	var carts []entity.DisplayCart

	for rows.Next() {
		var cart entity.DisplayCart
		err := rows.Scan(&cart.Id, &cart.Name, &cart.Description, &cart.Published)
		if err != nil {
			return []entity.DisplayCart{}, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

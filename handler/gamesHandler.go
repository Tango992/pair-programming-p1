package handler

import (
	"database/sql"
	"fmt"
	"pair-programming/entity"
)

func DisplayGames(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM games")
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

		fmt.Printf("[%v] - [%v] - [%v] - [%v]\n",games.Id, games.Name, games.Description, games.Published)
	}
	fmt.Println()
	return nil
}
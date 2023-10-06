package main

import (
	"fmt"
	"log"
	"pair-programming/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Koneksi berhasil")
}

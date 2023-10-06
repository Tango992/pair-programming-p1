package main

import (
	"fmt"
	"log"
	"pair-programming/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := config.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer config.DB.Close()

	fmt.Println("Koneksi berhasil")
}

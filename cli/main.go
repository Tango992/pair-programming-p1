package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pair-programming/config"
	"pair-programming/entity"
	"pair-programming/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Selamat datang di Steam CLI")
		fmt.Println("Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Masukkan pilihan (1/2/3): ")

		var option int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
		if err != nil {
			fmt.Println("Input anda bukan merupakan integer")
			continue
		}

		switch option {
		case 1:
			var username, password string
			fmt.Printf("\nREGISTER\n")
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)

			user := entity.User{
				Username: username,
				Password: password,
			}

			if err := handler.Register(user, db); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Registrasi berhasil!")
			}

		case 2:
			var username, password string
			fmt.Printf("\nLOGIN\n")
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			user, authenticated, _ := handler.Login(username, password, db)
			if authenticated {
				for {
					fmt.Printf("\nSelamat datang %v. Pilih opsi berikut:\n", user.Username)
					fmt.Println("1. Tampilkan semua game")
					fmt.Println("2. Tampilkan cart")
					fmt.Println("3. Tambah game ke cart")
					fmt.Println("4. Hapus game dari cart")
					fmt.Println("5. Exit")
					fmt.Print("Masukkan pilihan (1/2/3/4/5): ")

					var option int
					scanner.Scan()
					_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
					if err != nil {
						fmt.Println("Input anda bukan merupakan integer")
						continue
					}

					switch option {
					case 1:
						err := handler.DisplayGames(db)
						if err != nil {
							log.Fatal(err)
						}

					case 2:
						err := handler.DisplayCart(db, user)
						if err != nil {
							log.Fatal(err)
						}

					case 3:
						var gameOption int
						for {
							fmt.Print("Masukkan ID game: ")
							scanner.Scan()
							_, err := fmt.Sscanf(scanner.Text(), "%d", &gameOption)
							if err != nil {
								fmt.Println("Input bukan merupakan angka")
								continue
							}
							break
						}

						err = handler.AddCart(db, user, gameOption)
						if err != nil {
							log.Fatal(err)
						}

					case 4:
						var deleteGameOption int
						for {
							fmt.Print("Masukkan ID game yang ingin dihapus: ")
							scanner.Scan()
							_, err := fmt.Sscanf(scanner.Text(), "%d", &deleteGameOption)
							if err != nil {
								fmt.Println("Input bukan merupakan angka")
								continue
							}
							break
						}

						err = handler.DeleteCart(db, user, deleteGameOption)
						if err != nil {
							log.Fatal(err)
						}

					case 5:
						fmt.Println("Sampai jumpa!")
						os.Exit(0)

					default:
						fmt.Println("Input harus merupakan angka 1-5")
						os.Exit(1)
					}
				}
			} else {
				log.Fatal("Username / Password tidak sesuai")
			}

		case 3:
			fmt.Println("Sampai jumpa!")
			os.Exit(0)

		default:
			fmt.Println("Input harus merupakan angka 1-3!")
		}
	}
}

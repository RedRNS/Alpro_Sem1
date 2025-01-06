package main

import "fmt"

func main() {
	var user, pass, masukuser, masukpass string
	var percobaan int64
	user = "admin"
	pass = "admin"
	percobaan = 0
	fmt.Scan(&masukuser, &masukpass)

	for user != masukuser || pass != masukpass {
		percobaan++
		fmt.Scan(&masukuser, &masukpass)
	}
	fmt.Print(percobaan, "\nLogin Berhasil")

}

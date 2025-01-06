package main

import "fmt"

func main() {
	var ganjil bool
	var bil int
	fmt.Print("Bilangan : ")
	fmt.Scanln(&bil)
	ganjil = bil%2 != 0
	fmt.Println("Ganjil :", ganjil)
}

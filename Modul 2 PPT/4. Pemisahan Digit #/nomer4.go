package main

import "fmt"

func main() {
	var bil int
	print("Masukkan bilangan ratusan : ")
	fmt.Scanln(&bil)

	fmt.Printf("Hasil: %d %d %d", bil/100%10, bil/10%10, bil%10)
}

// Renisa Assyifa Putri [103112400123]

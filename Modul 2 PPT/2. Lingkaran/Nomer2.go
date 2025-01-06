package main

import "fmt"

func main() {
	var jari2, luas, keliling float64
	print("Masukkan jari-jari : ")
	fmt.Scanln(&jari2)

	luas = jari2 * jari2 * 3.14
	keliling = 2 * 3.14 * jari2

	fmt.Print("Luas: ", luas, "\nKeliling: ", keliling)
}

// Renisa Assyifa Putri [103112400123]

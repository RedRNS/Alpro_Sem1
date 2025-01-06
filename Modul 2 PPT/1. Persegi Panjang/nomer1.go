package main

import "fmt"

func main() {
	var panjang, lebar int
	print("Masukkan panjang dan lebar : ")
	fmt.Scanln(&panjang, &lebar)
	fmt.Printf("Keliling : %d \nLuas : %d\n", (2*panjang)+(2*lebar), (panjang * lebar))
}

// Renisa Assyifa Putri [103112400123]

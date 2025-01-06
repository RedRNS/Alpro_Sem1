package main

import "fmt"

func main() {
	var harga1, harga2, harga3 float64
	print("Masukkan 3 bilangan ratusan : ")
	fmt.Scanln(&harga1, &harga2, &harga3)

	harga1 = harga1 + harga1*5/100
	harga2 = harga2 + harga2*5/100
	harga3 = harga3 + harga3*5/100

	fmt.Println(harga1, harga2, harga3)
}

// Renisa Assyifa Putri [103112400123]

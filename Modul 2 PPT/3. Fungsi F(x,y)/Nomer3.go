package main

import "fmt"

func main() {
	var x, y, bil float64
	print("Masukkan 2 buah bilangan : ")
	fmt.Scanln(&x, &y)

	bil = 1/(3*x*x+10) + 10*y + 7
	fmt.Println("Hasil: ", bil)
}

// Renisa Assyifa Putri [103112400123]

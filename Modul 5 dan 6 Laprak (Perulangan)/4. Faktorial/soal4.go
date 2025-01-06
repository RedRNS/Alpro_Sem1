package main

import "fmt"

func main() {
	var n, i int
	var total int

	fmt.Print("Masukan bilangan bulat non negatif : ")
	fmt.Scanln(&n)
	total = 1

	for i = n; i >= 1; i -= 1 {
		total = total * i
	}

	fmt.Print("Hasil faktorial : ", total)
}

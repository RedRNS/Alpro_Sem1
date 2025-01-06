package main

import "fmt"

func main() {
	var n, i int
	var faktorial bool
	fmt.Scan(&n) // input bilangan bulat positif

	for i = 1; i <= n; i++ {
		faktorial = n%i == 0          // cek apakah i adalah faktor dari n
		fmt.Print("\n", i, faktorial) // output bilangan dan status faktornya
	}
}

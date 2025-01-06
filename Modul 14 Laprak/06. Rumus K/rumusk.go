package main

import "fmt"

func main() {
	var bil, atas, bawah, total float64

	fmt.Print("Nilai K : ")
	fmt.Scan(&bil)

	atas = (4*bil + 2) * (4*bil + 2)
	bawah = (4*bil + 1) * (4*bil + 3)
	total = atas / bawah

	fmt.Printf("Nilai f(K) = %0.10f", total)
}

package main

import "fmt"

func main() {
	var n, i, penyebut int64
	var total float64
	fmt.Scan(&n)
	penyebut = 1

	for i = 1; i <= n; i++ {
		if i%2 != 0 {
			total += 1 / float64(penyebut)
		} else {
			total -= 1 / float64(penyebut)
		}
		penyebut += 2
	}

	fmt.Printf("Hasil PI : %0.6f\n", 4*total)
}

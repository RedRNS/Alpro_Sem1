package main

import "fmt"

func main() {
	var n, angka, terbesar int
	fmt.Scan(&n)
	terbesar = 0

	for n > 0 {
		angka = n % 10
		if angka > terbesar {
			terbesar = angka
		}
		n = n / 10
	}
	fmt.Print(terbesar)
}

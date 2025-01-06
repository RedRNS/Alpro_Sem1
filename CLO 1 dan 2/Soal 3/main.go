package main

import "fmt"

func main() {
	var N, angka int
	fmt.Scanln(&N)

	totalawal := 0
	totalakhir := 0

	for i := 1; i <= N; i++ {
		fmt.Scan(&angka)
		totalawal += angka
		totalakhir += angka * 90 / 100
	}

	fmt.Print(totalawal, totalakhir)
}

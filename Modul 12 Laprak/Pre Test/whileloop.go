package main

import "fmt"

func main() {
	angka := 1
	total := 0

	for angka <= 10 {
		total += angka
		angka++
	}

	fmt.Println("Jumlah angka dari 1 hingga 10 adalah:", total)
}

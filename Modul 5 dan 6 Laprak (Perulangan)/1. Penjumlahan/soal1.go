package main

import "fmt"

func main() {
	var i, n, total int
	total = 0
	fmt.Print("Masukan Bilangan Bulat : ")
	fmt.Scan(&n)
	for i = 1; i <= n; i = i + 1 {
		total += i
	}
	fmt.Print("Hasil Penjumlahan : ", total)
}

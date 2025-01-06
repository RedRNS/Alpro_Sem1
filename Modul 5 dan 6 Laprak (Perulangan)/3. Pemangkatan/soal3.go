package main

import "fmt"

func main() {
	var angka, total, i, n int

	fmt.Print("Masukan angka dan pangkatnya : ")
	fmt.Scanln(&angka, &n)
	total = 1

	for i = 1; i <= n; i = i + 1 {
		total = total * angka
	}
	fmt.Print("Hasil pangkat : ", total)
}

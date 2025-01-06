package main

import "fmt"

func main() {
	var i, n, m int
	fmt.Print("Masukan 2 angka : ")
	fmt.Scan(&n, &m)

	for i = n; i <= m; i += 1 {
		fmt.Print(i, " ")
	}
}

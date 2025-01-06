package main

import "fmt"

func main() {
	var bilangan_n, angka, i int
	fmt.Scan(&bilangan_n)

	for i = 1; i <= bilangan_n; i++ {
		fmt.Scan(&angka)
		fmt.Print(angka*i, " ")
	}
}

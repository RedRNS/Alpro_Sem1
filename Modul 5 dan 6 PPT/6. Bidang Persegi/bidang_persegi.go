package main

import "fmt"

func main() {
	var n, i int
	var sisi, luas, keliling int

	fmt.Scanln(&n)

	for i = 1; i <= n; i += 1 {
		fmt.Scanln(&sisi)
		luas = sisi * sisi
		keliling = sisi * 4
		fmt.Println(luas, keliling)
	}
}

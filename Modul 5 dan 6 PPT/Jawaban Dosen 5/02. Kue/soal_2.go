package main

import "fmt"

func main() {
	var n, j, x, total_berat int
	fmt.Scan(&n)
	for j = 1; j <= n; j++ {
		fmt.Scan(&x)
		total_berat = x * 10
		fmt.Println(total_berat)
	}
}

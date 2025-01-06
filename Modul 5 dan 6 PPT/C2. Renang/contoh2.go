package main

import "fmt"

func main() {
	var n, j, x, totalJarak int
	fmt.Scanln(&n)

	for j = 1; j <= n; j += 1 {
		fmt.Scanln(&x)
		totalJarak = x * 5
		fmt.Println(totalJarak)
	}
}

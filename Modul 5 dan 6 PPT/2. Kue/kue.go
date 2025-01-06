package main

import "fmt"

func main() {
	var n, i, x int

	fmt.Scanln(&n)

	for i = 1; i <= n; i += 1 {
		fmt.Scan(&x)
		fmt.Print(x*i, " ")
	}
}

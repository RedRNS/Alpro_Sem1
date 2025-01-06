package main

import "fmt"

func main() {
	var n, daun, terbesar int
	fmt.Scan(&n)
	terbesar = 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&daun)
		if daun > terbesar {
			terbesar = daun
		}
	}
	fmt.Print(terbesar)
}

package main

import "fmt"

/* program dengan menggunakan while */

func main() {
	var N, iterasi int
	fmt.Scan(&N)
	iterasi = 0
	for iterasi != N {
		iterasi = iterasi + 1
		fmt.Print(iterasi, " ")

	}
}

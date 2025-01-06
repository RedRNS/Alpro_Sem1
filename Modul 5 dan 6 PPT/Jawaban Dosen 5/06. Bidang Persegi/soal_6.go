package main

import "fmt"

func main() {
	var n, s, i, luas, keliling int
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&s)
		luas = s * s
		keliling = 4 * s
		fmt.Println(luas, keliling)
	}
}

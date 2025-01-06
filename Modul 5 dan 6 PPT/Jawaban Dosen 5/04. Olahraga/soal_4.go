package main

import "fmt"

func main() {
	var n, steps, total, i, j int
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		total = 0
		for j = 0; j < 7; j++ {
			fmt.Scan(&steps)
			total += steps
		}
		fmt.Println(total)
	}
}

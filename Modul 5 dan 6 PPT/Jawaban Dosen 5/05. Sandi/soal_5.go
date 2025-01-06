package main

import "fmt"

func main() {
	var i, n, num, totalSum int
	totalSum = 0

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&num)

		totalSum += (num / 1000) + (num % 10)
	}

	fmt.Println(totalSum)
}

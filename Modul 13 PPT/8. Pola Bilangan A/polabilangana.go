package main

import "fmt"

func main() {
	var bil, i, j int64
	fmt.Scan(&bil)

	for i = 1; i <= bil; i++ {
		for j = 1; j <= bil; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

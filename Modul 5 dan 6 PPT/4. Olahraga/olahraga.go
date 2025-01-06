package main

import "fmt"

func main() {
	var orang, langkah, i, j, total int

	fmt.Scanln(&orang)
	total = 0

	for i = 1; i <= orang; i += 1 {
		for j = 1; j <= 7; j += 1 {
			fmt.Scan(&langkah)
			total = total + langkah
		}
		fmt.Println(total)
		total = 0
	}
}

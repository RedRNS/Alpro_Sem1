package main

import "fmt"

func main() {
	var bilangan, total int64
	fmt.Scan(&bilangan)
	total = 0

	for bilangan > 0 {
		fmt.Print(bilangan % 10)
		total = total + (bilangan % 10)
		bilangan = bilangan / 10
	}
	fmt.Print("\n", total)
}

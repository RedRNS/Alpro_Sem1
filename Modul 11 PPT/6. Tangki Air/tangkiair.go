package main

import "fmt"

func main() {
	var total, air, masuk int64
	fmt.Scan(&total)
	masuk = 0

	for masuk < total {
		fmt.Scan(&air)
		masuk += air
		if masuk < total {
			fmt.Print("False\n")
		}
	}
	fmt.Print("True")
}

package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	total := angka * angka

	for angka != 0 {
		fmt.Scan(&angka)
		total += angka * angka
	}

	fmt.Print(total)
}

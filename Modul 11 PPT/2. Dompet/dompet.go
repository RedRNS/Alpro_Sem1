package main

import "fmt"

func main() {
	var masuk, total int64
	fmt.Scan(&masuk)
	total = masuk

	for masuk != 0 {
		fmt.Scan(&masuk)
		total += masuk
	}
	fmt.Print(total)
}

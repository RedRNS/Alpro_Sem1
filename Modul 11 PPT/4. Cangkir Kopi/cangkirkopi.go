package main

import "fmt"

func main() {
	var gula, kopi, mgula, mkopi, cangkir int64
	fmt.Scan(&gula, &kopi, &mgula, &mkopi)
	cangkir = 0

	for gula > 0 && kopi > 0 {
		gula -= mgula
		kopi -= mkopi
		cangkir++
	}
	fmt.Print(cangkir)
}

package main

import "fmt"

func main() {
	var cari, angka, keep int
	var kebenaran bool
	kebenaran = false
	fmt.Scan(&cari, &angka)

	for angka > 0 {
		keep = angka % 10
		if cari == keep {
			kebenaran = true
			break
		}
		angka = angka / 10
	}
	fmt.Print(kebenaran)
}

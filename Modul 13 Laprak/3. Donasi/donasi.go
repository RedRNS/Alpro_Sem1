package main

import "fmt"

func main() {
	var target, bayar, jumlah, donatur int64
	jumlah = 0
	donatur = 0
	fmt.Scan(&target)

	for jumlah < target {
		fmt.Scan(&bayar)
		donatur++
		jumlah = jumlah + bayar
		fmt.Printf("Donatur %d : Menyumbang %d. Total terkumpul : %d\n", donatur, bayar, jumlah)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.", jumlah, donatur)
}

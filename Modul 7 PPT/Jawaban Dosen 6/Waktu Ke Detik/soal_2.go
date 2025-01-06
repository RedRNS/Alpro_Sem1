package main

import "fmt"

func main() {
	var totalDetik int

	fmt.Scan(&totalDetik)

	jam := totalDetik / 3600

	sisaDetik := totalDetik % 3600

	menit := sisaDetik / 60

	detik := sisaDetik % 60

	fmt.Printf("Waktu: %d jam, %d menit, %d detik\n", jam, menit, detik)
}

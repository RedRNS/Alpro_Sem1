package main

import "fmt"

func main() {
	var angka, jumlah int64
	jumlah = 0
	fmt.Scan(&angka)
	for done := false; !done; {
		angka = angka / 10
		jumlah++
		done = angka <= 0
	}
	fmt.Print(jumlah)
}

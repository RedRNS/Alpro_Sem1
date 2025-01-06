package main

import "fmt"

func main() {
	var N, angka int     //4321
	fmt.Scan(&N, &angka) // N = 4
	hasil := true
	for i := 0; i < N; i++ {
		digit1 := angka % 10               // 4321 --> digit1 = 1
		angka /= 10                        // angka = 5432
		digit2 := angka % 10               // digit2 = 2
		hasil = hasil && (digit1 < digit2) // 1<2
	}
	fmt.Print(hasil)
}

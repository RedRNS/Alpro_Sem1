package main

import "fmt"

func main() {
	var L, angkah int

	fmt.Scan(&L, &angkah)

	kebenaran := true

	for i := 1; i < L; i++ {
		nomor1 := angkah % 10
		angkah = angkah / 10
		nomor2 := angkah % 10
		kebenaran = kebenaran && (nomor1 < nomor2)
	}

	fmt.Print(kebenaran)
}

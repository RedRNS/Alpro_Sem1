package main

import "fmt"

type ukuran struct {
	panjang, lebar float64
	luas, keliling float64
}

func main() {
	var s ukuran

	fmt.Println("informasi Persegi Panjang: ")
	fmt.Print("Panjang : ")
	fmt.Scanln(&s.panjang)
	fmt.Print("Lebar : ")
	fmt.Scanln(&s.lebar)

	s.luas = s.lebar * s.panjang
	s.keliling = 2 * (s.lebar + s.panjang)

	fmt.Printf("Luas : %0.1f\n", s.luas)
	fmt.Printf("Keliling : %.1f", s.keliling)
}

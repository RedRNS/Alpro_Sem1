package main

import (
	"fmt"
	"math"
)

func main() {
	var i, n int
	var jari2, tinggi int
	var volume float64
	fmt.Print("Masukan Bilangan Bulat : ")
	fmt.Scanln(&n)
	for i = 1; i <= n; i = i + 1 {
		fmt.Print("Masukan Jari-jari dan Tinggi Kerucut : ")
		fmt.Scanln(&jari2, &tinggi)
		volume = math.Pi * 1 / 3 * float64(jari2) *
			float64(jari2) * float64(tinggi)
		fmt.Println("Volume Kerucut : ", volume)
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Masukkan nilai jari-jari : ")
	fmt.Scan(&x)
	r := math.Pi * x * x

	fmt.Printf("Nilai luas lingkaran adalah %.1f", r)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	var garis1, garis2, garis3 float64
	var garis_terpanjang float64

	fmt.Print("Koordinat titik 1 : ")
	fmt.Scanln(&x1, &y1)
	fmt.Print("Koordinat titik 2 : ")
	fmt.Scanln(&x2, &y2)
	fmt.Print("Koordinat titik 3 : ")
	fmt.Scanln(&x3, &y3)

	garis1 = math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
	garis2 = math.Sqrt(math.Pow(x2-x3, 2) + math.Pow(y2-y3, 2))
	garis3 = math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))

	garis_terpanjang = math.Max(garis1, math.Max(garis2, garis3))

	fmt.Printf("Panjang Sisi Terpanjang : %.2f", garis_terpanjang)
}

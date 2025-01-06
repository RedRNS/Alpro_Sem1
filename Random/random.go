package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)
	pi := 0.0
	for i := 0; i < n; i++ {
		suku := math.Pow(-1, float64(i)) / float64(2*i+1)
		pi += suku
	}
	pi *= 4
	fmt.Printf("Hasil PI: %.7f\n", pi)
	epsilon := 0.00001
	pi = 0.0
	sukuSebelumnya := 0.0
	i := 0
	for {
		suku := math.Pow(-1, float64(i)) / float64(2*i+1)
		pi += suku
		if math.Abs(suku-sukuSebelumnya) < epsilon {
			break
		}
		sukuSebelumnya = suku
		i++
	}
	fmt.Printf("Hasil PI: %.7f\n", pi*4)
	fmt.Printf("Pada i ke: %d\n", (i+1)*2)
}

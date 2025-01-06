package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x, y, A, B, C, D float64
	var n, i int64
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if x <= 0.5 && y <= 0.5 {
			A += 0.0001
		} else if x > 0.5 && y <= 0.5 {
			B += 0.0001
		} else if x > 0.5 && y > 0.5 {
			C += 0.0001
		} else if x <= 0.5 && y > 0.5 {
			D += 0.0001
		}
	}

	fmt.Printf("Curah hujan daerah A : %0.4f milimeter\n", A)
	fmt.Printf("Curah hujan daerah B : %0.4f milimeter\n", B)
	fmt.Printf("Curah hujan daerah C : %0.4f milimeter\n", C)
	fmt.Printf("Curah hujan daerah D : %0.4f milimeter\n", D)
}

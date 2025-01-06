package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, jumlahtoping, i int64
	var pusatX, pusatY, radius, radius2, pi float64

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	pusatX = 0.5
	pusatY = 0.5
	radius = 0.5
	radius2 = radius * radius
	jumlahtoping = 0

	rand.Seed(1) //supaya hasil randomnya konsisten

	for i = 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		dx := x - pusatX
		dy := y - pusatY

		if dx*dx+dy*dy <= radius2 {
			jumlahtoping++
		}
	}

	pi = 4.0 * float64(jumlahtoping) / float64(n)

	fmt.Println("Topping pada Pizza :", jumlahtoping)
	fmt.Printf("PI : %.10f\n", pi)
}

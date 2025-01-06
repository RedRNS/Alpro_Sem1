package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i, total_i, penyebut int64
	var total2, total1 float64
	fmt.Scan(&n)
	penyebut = 1

	for i = 1; i <= n; i++ {
		total1 = total2
		if i%2 != 0 {
			total2 += 1 / float64(penyebut)
		} else {
			total2 -= 1 / float64(penyebut)
		}

		penyebut += 2

		if math.Abs(total2-total1) <= 0.00001 {
			total_i = i
			break
		}
	}

	fmt.Printf("Hasil PI : %0.6f\n", 4*total1)
	fmt.Printf("Hasil PI : %0.6f\n", 4*total2)
	fmt.Println("Pada i ke :", total_i)
}

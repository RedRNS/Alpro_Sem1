package main

import (
	"fmt"
)

func main() {
	var banyak_hari, jam, totalJam, i int
	var rata_rata float64
	fmt.Scan(&banyak_hari)

	for i = 0; i < banyak_hari; i++ {
		fmt.Scan(&jam)
		totalJam += jam
	}

	rata_rata = float64(totalJam) / float64(banyak_hari)

	fmt.Printf("%.3f", rata_rata)
}

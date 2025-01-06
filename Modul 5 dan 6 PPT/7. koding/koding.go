package main

import "fmt"

func main() {
	var i int
	var jam, orang, total int
	var rata2 float64

	fmt.Scanln(&orang)
	total = 0

	for i = 1; i <= orang; i += 1 {
		fmt.Scanln(&jam)
		total += jam
	}

	rata2 = float64(total) / float64(orang)

	fmt.Printf("%0.3f", rata2)
}

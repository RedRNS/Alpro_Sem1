package main

import "fmt"

func main() {
	var bil, atas, bawah, i float64
	var total = 1.0

	fmt.Print("Nilai K : ")
	fmt.Scan(&bil)

	for i = 0; i <= bil; i++ {
		atas = (4*i + 2) * (4*i + 2)
		bawah = (4*i + 1) * (4*i + 3)
		total = total * (atas / bawah)
	}

	fmt.Printf("Nilai akar 2 = %0.10f", total)
}

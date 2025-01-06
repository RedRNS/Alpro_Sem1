package main

import "fmt"

func main() {
	var banyak, i int64
	var urutan, bunga string
	fmt.Print("N : ")
	fmt.Scanln(&banyak)

	for i = 1; i <= banyak; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		if i == 1 {
			urutan += bunga
		} else {
			urutan = urutan + " - " + bunga
		}

	}
	fmt.Print("Pita : ", urutan)
}

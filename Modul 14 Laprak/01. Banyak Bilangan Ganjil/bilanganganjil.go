package main

import "fmt"

func main() {
	var bil, total, i int64
	fmt.Scan(&bil)

	for i = 1; i <= bil; i++ {
		if i%2 != 0 {
			total++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", total)
}

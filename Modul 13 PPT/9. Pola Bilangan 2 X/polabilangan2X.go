package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan angka (harus ganjil): ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Masukkan angka ganjil untuk hasil X yang simetris.")
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				fmt.Printf("%d", i+1)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

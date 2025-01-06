package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // Membaca bilangan bulat positif

	// Menetapkan nilai faktorialnya 1 jika n adalah 0
	faktorial := 1

	// Mengalikan setiap bilangan dari 1 hingga n
	for i := 1; i <= n; i++ {
		faktorial *= i
	}

	fmt.Print(faktorial)
}

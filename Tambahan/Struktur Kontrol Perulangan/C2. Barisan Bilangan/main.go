package main

import "fmt"

/* salah satu solusi menggunakan repeat */

func main() {
	var N, iterasi int
	var berhenti bool
	fmt.Scan(&N)
	iterasi = 0
	berhenti = false
	for !berhenti {
		iterasi = iterasi + 1
		fmt.Print(iterasi, " ")
		berhenti = iterasi == N

	}
}

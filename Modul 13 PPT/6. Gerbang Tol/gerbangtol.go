package main

import "fmt"

func main() {
	var huruf string
	var JA, JB, JC int = 0, 0, 0
	fmt.Scan(&huruf)

	for huruf == "A" || huruf == "B" || huruf == "C" {
		if huruf == "A" {
			JA++
		} else if huruf == "B" {
			JB++
		} else {
			JC++
		}
		fmt.Scan(&huruf)
	}

	fmt.Println("Tipe A = ", JA)
	fmt.Println("Tipe B = ", JB)
	fmt.Println("Tipe C = ", JC)
}

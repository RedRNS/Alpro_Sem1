package main

import "fmt"

func main() {
	var bil int

	fmt.Scan(&bil)

	if bil%2 == 0 && bil < 0 {
		fmt.Print("genap negatif")
	} else {
		fmt.Print("bukan")
	}
}

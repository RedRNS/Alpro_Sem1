package main

import "fmt"

func main() {
	var bil, jumlah int
	fmt.Scan(&bil)
	jumlah = 0
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			jumlah++
		}
	}

	if jumlah == 2 {
		fmt.Print("True")
	} else if bil == 1 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}

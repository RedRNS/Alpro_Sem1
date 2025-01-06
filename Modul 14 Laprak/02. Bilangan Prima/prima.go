package main

import "fmt"

func main() {
	var bil, i int64
	var prima string
	prima = "prima"
	fmt.Scan(&bil)

	for i = 2; i < bil; i++ {
		if bil%i == 0 {
			prima = "bukan prima"
			break
		}
	}
	fmt.Print(prima)
}

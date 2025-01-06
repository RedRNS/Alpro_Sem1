package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukan : ")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		for j := 1; j <= a; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

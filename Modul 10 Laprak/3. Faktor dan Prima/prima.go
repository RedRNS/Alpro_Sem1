package main

import "fmt"

func main() {
	var bil, i, faktor int
	var prima bool = true

	fmt.Scan(&bil)
	fmt.Println("Bilangan : ", bil)
	fmt.Print("Faktor : ")

	if bil > 1 {
		for i = 1; i <= bil; i++ {
			if bil%i == 0 {
				fmt.Print(i, " ")
				faktor++
			}
		}
	}

	if faktor == 2 {
		fmt.Println("\nPrima : ", prima)
	} else {
		fmt.Println("\nPrima : ", !prima)
	}
}

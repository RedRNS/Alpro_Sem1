package main

import "fmt"

func main() {
	var angka, i int
	var kalimat string

	fmt.Scanln(&angka)
	fmt.Scanln(&kalimat)

	for i = 1; i <= angka; i += 1 {
		fmt.Println(kalimat)
	}
}

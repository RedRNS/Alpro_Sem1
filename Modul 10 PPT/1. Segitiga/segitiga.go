package main

import "fmt"

func main() {
	var sisi1, sisi2, sisi3 int
	fmt.Scan(&sisi1, &sisi2, &sisi3)

	if sisi1 == sisi2 && sisi2 == sisi3 {
		fmt.Print("Segitiga Sama Sisi")
	} else if sisi1 == sisi2 || sisi2 == sisi3 || sisi1 == sisi3 {
		fmt.Print("Segitiga Sama Kaki")
	} else {
		fmt.Print("Segitiga Sembarang")
	}
}

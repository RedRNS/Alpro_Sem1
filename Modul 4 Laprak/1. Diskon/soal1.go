package main

import "fmt"

func main() {
	var awal, diskon, akhir int
	fmt.Print("Masukan nilai awal dan diskon : ")
	fmt.Scanln(&awal, &diskon)

	akhir = awal * (100 - diskon) / 100

	print("Total belanja akhir : ", akhir)
}

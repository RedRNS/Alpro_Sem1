package main

import "fmt"

func main() {
	var F, C float64
	fmt.Print("Masukan Suhu dalam satuan Fahrenheit : ")
	fmt.Scan(&F)
	C = 5.0 / 9.0 * (F - 32)
	fmt.Print("Hasil konversi ke Celcius : ", C)
}

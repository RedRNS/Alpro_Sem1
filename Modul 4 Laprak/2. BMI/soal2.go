package main

import "fmt"

func main() {
	var BMI, tinggi, berat float64

	fmt.Print("Masukan nilai BMI dan tinggi badan : ")
	fmt.Scanln(&BMI, &tinggi)

	berat = BMI * tinggi * tinggi

	fmt.Printf("Berat badan : %.0f", berat)
}

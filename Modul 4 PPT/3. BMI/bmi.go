package main

import "fmt"

type BMI struct {
	nama          string
	berat, tinggi float64
	total         float64
}

func main() {
	var bmi BMI

	fmt.Print("Masukan nama, berat, dan tinggi badan : ")
	fmt.Scanln(&bmi.nama, &bmi.berat, &bmi.tinggi)

	bmi.total = bmi.berat / (bmi.tinggi * bmi.tinggi)

	fmt.Println("\nInformasi BMI : ")
	fmt.Println("Nama : ", bmi.nama)
	fmt.Printf("Berat : %.1f kg\n", bmi.berat)
	fmt.Printf("Tinggi : %.2f m\n", bmi.tinggi)
	fmt.Printf("BMI : %.2f\n", bmi.total)
}

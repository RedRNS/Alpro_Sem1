package main

import "fmt"

func main() {
	var C, F, R, K float64
	fmt.Print("Temperatur Celcius : ")
	fmt.Scanln(&C)
	F = C*9/5 + 32
	R = C * 4 / 5
	K = (F + 459.67) * 5 / 9
	fmt.Println("Derajat Reamur :", R)
	fmt.Println("Derajat Fahrenheit :", F)
	fmt.Printf("Derajat Kelvin : %.0f", K)
}

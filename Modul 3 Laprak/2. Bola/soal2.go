package main

import "fmt"

func main() {
	var jari, luaskulit, volume float64
	fmt.Print("Jejari : ")
	fmt.Scanln(&jari)
	pi := 3.1415926535
	luaskulit = 4.0 * pi * jari * jari
	volume = (4.0 / 3.0) * pi * jari * jari * jari
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f", jari, volume, luaskulit)
}

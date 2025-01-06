package main

import "fmt"

const PI float64 = 3.14

type tabung struct {
	tinggi, jari2         int
	luas, volume          float64
	luasAlas, luasDinding float64
}

func main() {
	var t tabung
	fmt.Print("Masukan jari-jari dan tinggi : ")
	fmt.Scanln(&t.jari2, &t.tinggi)
	t.luasAlas = PI * (float64(t.jari2) * float64(t.jari2))
	t.luasDinding = float64(t.tinggi) * (2 * PI * float64(t.jari2))
	t.luas = 2*t.luasAlas + t.luasDinding
	t.volume = t.luasAlas * float64(t.tinggi)
	fmt.Print(t.luas, t.volume)
}

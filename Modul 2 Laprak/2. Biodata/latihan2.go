package main

import "fmt"

func main() {
	var x, y, z string
	fmt.Println("Silahkan masukkan nama, NIM dan kelas!")
	fmt.Scanln(&x, &y, &z)
	fmt.Printf("Halo, perkenalkan semuanya, nama saya %s, "+
		"NIM saya %s dan saya dari kelas %s. "+
		"Salam Kenal Semuanya!", x, y, z)
}

package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Masukan sebuah bilangan bulat : ")
	fmt.Scanln(&x)
	x = 2/(x+5) + 5
	fmt.Println("Hasil : ", x)
}

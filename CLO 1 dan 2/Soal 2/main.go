package main

import "fmt"

func main() {
	var a, b, c, d, e, angka int

	fmt.Scan(&angka)

	a = angka % 10
	b = angka / 10 % 10
	c = angka / 100 % 10
	d = angka / 1000 % 10
	e = angka / 10000

	fmt.Printf("%d%d%d%d%d", a, b, c, d, e)
}

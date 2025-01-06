package main

import "fmt"

func main() {
	var a, b, bilangan int64
	fmt.Scan(&bilangan)
	kebenaran := true

	for bilangan > 0 {
		a = bilangan % 10
		b = bilangan / 10 % 10
		hasil := a - b
		if hasil < 0 {
			hasil = hasil * -1
		}
		if hasil != 1 {
			kebenaran = false
			break
		}
		bilangan = bilangan / 10
	}
	fmt.Print(kebenaran)
}

package main

import "fmt"

func main() {
	var x1, y1, a1, x2, y2, a2, hasil1, hasil2 float64

	fmt.Scanln(&x1, &y1, &a1)
	fmt.Scanln(&x2, &y2, &a2)

	hasil1 = (1.0-a1)*(3.0*x1/(2.0*y1)) + (a1*4.0*y1)/(5.0*x1)
	hasil2 = (1.0-a2)*(3.0*x2/(2.0*y2)) + (a2*4.0*y2)/(5.0*x2)

	fmt.Println(hasil1)
	fmt.Println(hasil2)
}

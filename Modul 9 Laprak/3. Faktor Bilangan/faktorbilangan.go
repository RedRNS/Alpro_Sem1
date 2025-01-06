package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	if y%x == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if x%y == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

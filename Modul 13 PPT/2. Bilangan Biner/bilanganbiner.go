package main

import "fmt"

func main() {
	var num int
	var binary string
	fmt.Scan(&num)

	for num > 0 {
		remainder := num % 2
		if remainder == 1 {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
		num = num / 2
	}
	
	fmt.Println(binary)
}

package main

import "fmt"

func main() {
	var n, i int
	var bilangan, total int
	var depan, belakang int

	fmt.Scanln(&n)
	total = 0

	for i = 1; i <= n; i += 1 {
		fmt.Scanln(&bilangan)
		depan = bilangan / 1000 % 10
		belakang = bilangan % 10
		total += depan + belakang
	}
	fmt.Println(total)
}

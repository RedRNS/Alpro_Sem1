package main

import "fmt"

func main() {
	var i, N, d int
	var s bool

	fmt.Scanln(&N)

	for i = 1; i <= N; i += 1 {
		s = 5%i == 0
		d = i
		println(d, s)
	}
}

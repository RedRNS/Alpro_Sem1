package main

import "fmt"

func main() {
	var manusia, motor int

	fmt.Scan(&manusia)

	if manusia%2 == 0 {
		motor = manusia / 2
	} else {
		motor = (manusia / 2) + 1
	}

	fmt.Println(motor)
}

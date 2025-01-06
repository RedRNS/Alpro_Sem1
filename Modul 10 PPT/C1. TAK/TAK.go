package main

import "fmt"

func main() {
	var ikk float64
	fmt.Scan(&ikk)
	if ikk < 2.00 {
		fmt.Println("Poor")
	} else if 2.00 <= ikk && ikk <= 2.75 {
		fmt.Println("Fair")
	} else if 2.76 <= ikk && ikk <= 3.00 {
		fmt.Println("Sastifactory")
	} else if 3.01 <= ikk && ikk <= 3.50 {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Excellent")
	}
}

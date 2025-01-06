package main

import "fmt"

func main() {
	var T1, T2, T3, T4, T5 int
	fmt.Scan(&T1, &T2, &T3, &T4, &T5)

	if T1 <= T2 && T2 <= T3 && T3 <= T4 && T4 <= T5 {
		fmt.Print("Stabil Naik")
	} else if T1 >= T2 && T2 >= T3 && T3 >= T4 && T4 >= T5 {
		fmt.Print("Stabil Turun")
	} else {
		fmt.Print("Tidak Stabil")
	}
}

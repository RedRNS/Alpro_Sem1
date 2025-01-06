package main

import "fmt"

func main() {
	var nilai, TAK float64
	var status string
	fmt.Scan(&nilai)

	if nilai > 3.0 {
		fmt.Scan(&TAK)
		if TAK < 2.00 {
			status = "Poor"
		} else if 2.00 <= TAK && TAK <= 2.75 {
			status = "Fair"
		} else if 2.76 <= TAK && TAK <= 3.00 {
			status = "Satisfactory"
		} else if 3.01 <= TAK && TAK <= 3.50 {
			status = "Very Good"
		} else {
			status = "Excellent"
		}
		fmt.Printf("Cumlaude dengan IPK %0.2f dan status predikat TAK adalah %s", nilai, status)
	} else {
		fmt.Print("Tidak Cumlaude")
	}
}

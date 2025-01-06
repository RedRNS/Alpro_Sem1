package main

import "fmt"

func main() {
	var nilai, TAK float64
	fmt.Scan(&nilai)

	if nilai > 3.0 {
		fmt.Scan(&TAK)
		if TAK < 2.00 {
			fmt.Printf("Cumlaude dengan IPK %f dan status predikat TAK adalah Poor", nilai)
		} else if 2.00 <= TAK && TAK <= 2.75 {
			fmt.Printf("Cumlaude dengan IPK %f dan status predikat TAK adalah Fair", nilai)
		} else if 2.76 <= TAK && TAK <= 3.00 {
			fmt.Printf("Cumlaude dengan IPK %f dan status predikat TAK adalah Satisfactory", nilai)
		} else if 3.01 <= TAK && TAK <= 3.50 {
			fmt.Printf("Cumlaude dengan IPK %f dan status predikat TAK adalah Very Good", nilai)
		} else {
			fmt.Printf("Cumlaude dengan IPK %f dan status predikat TAK adalah Excellent", nilai)
		}
	} else {
		fmt.Println("Tidak Cumlaude")
	}

}

//diatas 3.0 cumlaude --> punya hak akses input Tidak
//dibawah 3.0 tidak cumlaude --> tidak punya hak akses

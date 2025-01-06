package main

import "fmt"

func main() {
	var semester, skor int
	var lulus bool

	fmt.Print("Jumlah semester dan skor EPrT : ")
	fmt.Scanln(&semester, &skor)

	lulus = semester <= 8 && skor >= 500

	fmt.Println("Cumlaude :", lulus)
}

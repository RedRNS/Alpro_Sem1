package main

import "fmt"

func main() {
	var x, kata, ada, posisi string
	var min bool
	var n, i, total int64

	ada = "Tidak terdapat"

	fmt.Scanln(&x)
	fmt.Scanln(&n)

	for i = 1; i <= n; i++ {
		fmt.Scan(&kata)
		if x == kata {
			ada = "Terdapat"
			if posisi == "" {
				posisi = fmt.Sprintf("%d", i)
			} else {
				posisi = posisi + fmt.Sprintf(", %d", i)
			}
			total++
			if total >= 2 {
				min = true
			}
		}
	}

	fmt.Printf("\nJawaban a : %s string '%s' dalam %d kumpulan data.\n", ada, x, n)
	if ada == "Terdapat" {
		fmt.Printf("Jawaban b : Pada posisi ke-%s string '%s' ditemukan\n", posisi, x)
	} else {
		fmt.Printf("Jawaban b : String '%s' tidak dapat ditemukan\n", x)
	}
	fmt.Printf("Jawaban c : Terdapat %d buah string '%s' dalam kumpulan data.\n", total, x)
	if min == true {
		fmt.Printf("Jawaban d : Ada sedikitnya dua string '%s' dalam kumpulan data.\n", x)
	} else {
		fmt.Printf("Jawaban d : Terdapat kurang dari dua string '%s' dalam kumpulan data.\n", x)
	}
}

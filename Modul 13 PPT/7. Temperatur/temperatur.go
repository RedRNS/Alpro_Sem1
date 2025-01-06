package main

import "fmt"

func main() {
	var (
		bil, iterasi, rata2 int64
		tertinggi           int64 = 0
		terendah            int64 = 0
		nol                 int64 = 0
	)

	for {
		fmt.Scan(&bil)
		if bil < terendah {
			terendah = bil
		}

		if bil > tertinggi {
			tertinggi = bil
		}

		if bil == 0 {
			nol++
		} else {
			nol = 0
		}

		if nol == 2 {
			break
		}

		rata2 += bil
		iterasi++
	}
	fmt.Println("Tertinggi : ", tertinggi)
	fmt.Println("Terendah : ", terendah)
	fmt.Println("Rata-Rata : ", float64(rata2)/float64(iterasi))
}

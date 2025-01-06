package main

import "fmt"

func main() {
	var kiri, kanan float64 = 0, 0

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&kiri, &kanan)
		if (kanan+kiri > 150.0) || (kanan < 0) || (kiri < 0) {
			break
		}
		if kiri-kanan >= 9 || kanan-kiri > 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng : true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng : false")
		}
	}
	fmt.Print("Proses selesai.")
}

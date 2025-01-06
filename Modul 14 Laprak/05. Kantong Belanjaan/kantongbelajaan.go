package main

import "fmt"

func main() {
	var kiri, kanan float64 = 0, 0

	for kiri < 9.0 && kanan < 9.0 {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&kiri, &kanan)
	}
	fmt.Print("Proses selesai.")
}

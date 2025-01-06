package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	var tarif int
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)
	switch kendaraan {
	case "Motor":
		switch {
		case durasi >= 1 && durasi <= 2:
			tarif = 7000
		case durasi > 2:
			tarif = 9000
		}
	case "Mobil":
		switch {
		case durasi >= 1 && durasi <= 2:
			tarif = 7000
		case durasi > 2:
			tarif = 7000
		}
	case "Truk":
		switch {
		case durasi >= 1 && durasi <= 2:
			tarif = 7000
		case durasi > 2:
			tarif = 7000
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid.")
	}
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}

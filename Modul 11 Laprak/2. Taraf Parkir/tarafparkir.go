package main

import "fmt"

func main() {
	var kendaraan string
	var tarif float64
	var durasi float64

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch kendaraan {
	case "Motor":
		tarif = durasi * 2000
	case "Mobil":
		tarif = durasi * 5000
	case "Truk":
		tarif = durasi * 8000
	}

	fmt.Printf("Tarif Parkir: Rp %.0f\n", tarif)
}

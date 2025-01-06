package main

import "fmt"

type karyawan struct {
	nama                            string
	gaji_pokok, tunjangan, potongan float64
	total_gaji                      float64
}

func main() {
	var data karyawan

	fmt.Print("Masukan nama : ")
	fmt.Scanln(&data.nama)
	fmt.Print("Gaji Pokok : ")
	fmt.Scanln(&data.gaji_pokok)
	fmt.Print("Tunjangan : ")
	fmt.Scanln(&data.tunjangan)
	fmt.Print("Potongan : ")
	fmt.Scanln(&data.potongan)

	data.total_gaji = data.gaji_pokok + data.tunjangan - data.potongan

	fmt.Println("\nInformasi Karyawan : ")
	fmt.Println("Nama : ", data.nama)
	fmt.Printf("Gaji Pokok : Rp %.2f\n", data.gaji_pokok)
	fmt.Printf("Potongan : Rp %.2f\n", data.potongan)
	fmt.Printf("Total Gaji : Rp %.2f\n", data.total_gaji)
}

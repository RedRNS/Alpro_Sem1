package main

import "fmt"

type transaksi struct {
	nama         string
	jumlah       int
	harga, total float64
}

func main() {
	var trans transaksi

	fmt.Println("Informasi Transaksi : ")
	fmt.Print("Nama barang : ")
	fmt.Scanln(&trans.nama)
	fmt.Print("Jumlah : ")
	fmt.Scanln(&trans.jumlah)
	fmt.Print("Harga Satuan : Rp ")
	fmt.Scanln(&trans.harga)

	trans.total = float64(trans.jumlah) * trans.harga
	fmt.Printf("Total Harga : Rp %.2f", trans.total)
}

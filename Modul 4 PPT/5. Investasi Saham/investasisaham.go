package main

import (
	"fmt"
	"math"
)

type invest struct {
	harga_beli           float64
	harga_jual           float64
	jumlah_saham         float64
	total_investasi_awal float64
	total_penjualan      float64
	keuntungan_kotor     float64
	biaya_transaksi      float64
	pajak_keuntungan     float64
	keuntungan_bersih    float64
}

func main() {
	var inv invest

	fmt.Println("\nInformasi Investasi Saham : ")
	fmt.Print("Harga beli : Rp ")
	fmt.Scanln(&inv.harga_beli)
	fmt.Print("Harga jual : Rp ")
	fmt.Scanln(&inv.harga_jual)
	fmt.Print("Jumlah saham : ")
	fmt.Scanln(&inv.jumlah_saham)

	inv.total_investasi_awal = inv.harga_beli * inv.jumlah_saham
	inv.total_penjualan = inv.harga_jual * inv.jumlah_saham
	inv.keuntungan_kotor = inv.total_penjualan - inv.total_investasi_awal
	inv.biaya_transaksi = inv.total_penjualan * 0.2 / 100
	inv.pajak_keuntungan = math.Max(0, inv.keuntungan_kotor*10/100)
	inv.keuntungan_bersih = inv.keuntungan_kotor - inv.biaya_transaksi - inv.pajak_keuntungan

	fmt.Printf("Total Investasi Awal : Rp %.2f\n", inv.total_investasi_awal)
	fmt.Printf("Total Penjualan : Rp %.2f\n", inv.total_penjualan)
	fmt.Printf("Keuntungan Kotor : Rp %.2f\n", inv.keuntungan_kotor)
	fmt.Printf("Biaya Transaksi : Rp %.2f\n", inv.biaya_transaksi)
	fmt.Printf("Pajak Keuntungan : Rp %.2f\n", inv.pajak_keuntungan)
	fmt.Printf("Keuntungan Bersih : Rp %.2f\n", inv.keuntungan_bersih)
}

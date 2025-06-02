// Fungsi untuk menghitung total pembayaran setelah diskon
package main

import (
	"fmt"
)

func hitungDiskon(totalBelanja float64) float64 {
	var diskon float64
	if totalBelanja > 1000 {
		diskon = 0.2
	} else if totalBelanja >= 500 {
		diskon = 0.15
	} else {
		diskon = 0.05
	}
	return totalBelanja * (1 - diskon)
}

func main1() {
	var totalBelanja float64
	fmt.Print("Masukkan total belanja Mita: ")
	fmt.Scan(&totalBelanja)

	jumlahBayar := hitungDiskon(totalBelanja)
	fmt.Printf("Jumlah yang harus dibayar setelah diskon adalah: %.0f\n", jumlahBayar)
}

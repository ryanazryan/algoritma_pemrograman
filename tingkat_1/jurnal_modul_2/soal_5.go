package main

import (
	"fmt"
)

func main5() {
	var jumlahPermen, hargaPermen int
	fmt.Print("Masukkan jumlah permen yang dibeli dan harga satuannya: ")
	fmt.Scan(&jumlahPermen, &hargaPermen)

	var jumlahPensil, hargaPensil int
	fmt.Print("Masukkan jumlah pensil yang dibeli dan harga satuannya: ")
	fmt.Scan(&jumlahPensil, &hargaPensil)

	totalBayarPermen := jumlahPermen * hargaPermen
	bonusPermen := (jumlahPermen * 2) * hargaPermen

	totalBayarPensil := jumlahPensil * hargaPensil
	bonusPensil := (jumlahPensil * 2) * hargaPensil

	totalBayar := totalBayarPermen + totalBayarPensil
	totalPromo := bonusPermen + bonusPensil

	fmt.Printf("Total yang harus dibayar: %d\n", totalBayar)
	fmt.Printf("Total promo yang diperoleh: %d\n", totalPromo)

	// Penjelasan singkat
	// hitung total yang harus dibayar
	// total bayar = (jumlah bayar x harga permen) + (jumlah pensil x harga pensil)
	// lalu hitung value promo
	// total promo = (2 x jumlah permen x harga permen) + (2 x jumlah pensil x harga pensil)
}

package main

import (
	"fmt"
)

func konversiKeWaktuMars(detik int) (int, int, int) {
	jam := detik / 4500
	detikSisa := detik % 4500
	menit := detikSisa / 75
	detikAkhir := detikSisa % 75
	return jam, menit, detikAkhir
}

func main3() {
	var waktu1, waktu2, waktu3 int

	fmt.Print("Masukkan waktu pertama (dalam detik): ")
	fmt.Scan(&waktu1)
	fmt.Print("Masukkan waktu kedua (dalam detik): ")
	fmt.Scan(&waktu2)
	fmt.Print("Masukkan waktu ketiga (dalam detik): ")
	fmt.Scan(&waktu3)

	rataRataDetik := (waktu1 + waktu2 + waktu3) / 3

	jam, menit, detik := konversiKeWaktuMars(rataRataDetik)

	fmt.Printf("Hasil kalibrasi waktu %d jam, %d menit dan %d detik\n", jam, menit, detik)
}
package main

import (
	"fmt"
)

func main1() {
	var nomor int
	fmt.Print("Masukkan nomor polisi (hingga 5 digit): ")
	fmt.Scan(&nomor)

	hasil := jumlahDigitTunggal(nomor)

	fmt.Printf("Hasil penjumlahan digit: %d\n", hasil)
}

func jumlahDigitTunggal(n int) int {
	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}
	return n
}

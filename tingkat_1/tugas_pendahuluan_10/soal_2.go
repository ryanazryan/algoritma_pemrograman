package main

import (
	"fmt"
)

func cetakNominal(angka int) string {
	switch {
	case angka >= 1 && angka < 10:
		return "Satuan"
	case angka >= 10 && angka < 100:
		return "Puluhan"
	case angka >= 100 && angka < 1000:
		return "Ratusan"
	case angka >= 1000 && angka < 10000:
		return "Ribuan"
	case angka >= 10000 && angka < 100000:
		return "Puluhan ribu"
	case angka >= 100000 && angka <= 999999:
		return "Ratusan ribu"
	default:
		return "Nominal tidak valid"
	}
}

func main2() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	hasil := cetakNominal(angka)
	fmt.Printf("Nominal: %s\n", hasil)
}

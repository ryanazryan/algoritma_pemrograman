package main

import (
	"fmt"
)

func main4() {
	const hargaA = 85000
	const hargaB = 120000
	const hargaC = 250000

	var jumlahHari int
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&jumlahHari)

	totalPendapatan := 0

	for i := 1; i <= jumlahHari; i++ {
		var pesananA, pesananB, pesananC int
		fmt.Printf("Masukkan jumlah pesanan untuk hari ke-%d (A B C): ", i)
		fmt.Scan(&pesananA, &pesananB, &pesananC)

		pendapatanHari := (pesananA * hargaA) + (pesananB * hargaB) + (pesananC * hargaC)
		totalPendapatan += pendapatanHari
	}

	fmt.Printf("Total Pendapatan: %d\n", totalPendapatan)
}

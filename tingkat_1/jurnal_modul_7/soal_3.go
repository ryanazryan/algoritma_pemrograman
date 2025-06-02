package main

import (
	"fmt"
)

func main3() {
	var jumlahHari int
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&jumlahHari)

	totalMenit := 0

	for i := 1; i <= jumlahHari; i++ {
		var durasi int
		fmt.Printf("Masukkan durasi (dalam menit) untuk hari ke-%d: ", i)
		fmt.Scan(&durasi)
		totalMenit += durasi
	}

	totalJam := totalMenit / 60
	sisaMenit := totalMenit % 60

	fmt.Printf("Total waktu yang diperlukan adalah %d jam %d menit\n", totalJam, sisaMenit)
}

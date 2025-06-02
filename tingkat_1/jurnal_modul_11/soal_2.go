package main

import "fmt"

func main2() {
	var bulan string
	fmt.Println("Masukkan nama bulan (contoh: Januari, Februari, dll):")
	fmt.Scanln(&bulan)

	switch bulan {
	case "Desember", "Januari", "Februari":
		fmt.Println("Musim: Musim Dingin")
		fmt.Println("Suhu Rata-rata: -10 hingga 0 derajat C")
	case "Maret", "April", "Mei":
		fmt.Println("Musim: Musim Semi")
		fmt.Println("Suhu Rata-rata: 5 hingga 15 derajat C")
	case "Juni", "Juli", "Agustus":
		fmt.Println("Musim: Musim Panas")
		fmt.Println("Suhu Rata-rata: 20 hingga 35 derajat C")
	case "September", "Oktober", "November":
		fmt.Println("Musim: Musim Gugur")
		fmt.Println("Suhu Rata-rata: 10 hingga 20 derajat C")
	default:
		fmt.Println("Bulan tidak valid. Masukkan nama bulan dengan benar.")
	}
}

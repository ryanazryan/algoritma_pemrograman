package main

import (
	"fmt"
)

func main2() {
	var nilaiUjian, presensi, nilaiRemedial float64

	fmt.Print("Masukkan nilai ujian: ")
	fmt.Scan(&nilaiUjian)
	fmt.Print("Masukkan presensi: ")
	fmt.Scan(&presensi)

	if nilaiUjian < 50.01 && presensi >= 75 {
		fmt.Print("Masukkan nilai remedial: ")
		fmt.Scan(&nilaiRemedial)

		if nilaiRemedial > 55 {
			nilaiRemedial = 55
		}
		fmt.Printf("Nilai mahasiswa adalah %.0f\n", nilaiRemedial)
	} else {
		fmt.Printf("Nilai mahasiswa adalah %.0f\n", nilaiUjian)
	}
}

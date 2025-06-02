package main

import (
	"fmt"
)

func main2() {
	var lulus, tidakLulus int

	fmt.Print("Masukkan jumlah mahasiswa yang lulus: ")
	fmt.Scan(&lulus)
	fmt.Print("Masukkan jumlah mahasiswa yang tidak lulus: ")
	fmt.Scan(&tidakLulus)

	total := float64(lulus + tidakLulus)

	persenLulus := float64(lulus) / total
	persenTidakLulus := float64(tidakLulus) / total

	fmt.Printf("Kelulusan: %.2f Ketidaklulusan: %.2f\n", persenLulus, persenTidakLulus)
}

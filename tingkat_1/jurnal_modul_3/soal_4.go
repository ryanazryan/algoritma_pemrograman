package main

import (
	"fmt"
)

func hitungBeratIdeal(tinggi int, faktor float64) float64 {
	return float64(tinggi-100) - faktor*float64(tinggi-100)
}

func main4() {
	var tinggi int

	fmt.Print("Masukkan tinggi badan (dalam cm): ")
	fmt.Scan(&tinggi)

	beratIdealPangeran := hitungBeratIdeal(tinggi, 0.10)
	beratIdealPutri := hitungBeratIdeal(tinggi, 0.15)

	fmt.Printf("Berat badan ideal untuk Pangeran adalah %.1f kg dan Tuan Putri adalah %.1f kg\n", beratIdealPangeran, beratIdealPutri)
}
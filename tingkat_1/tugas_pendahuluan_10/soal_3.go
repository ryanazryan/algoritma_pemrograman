package main

import (
	"fmt"
)

func hitungBMI(berat int, tinggi int) string {
	tinggiMeter := float64(tinggi) / 100
	bmi := float64(berat) / (tinggiMeter * tinggiMeter)

	switch {
	case bmi < 18.5:
		return "Berat badan kurang"
	case bmi >= 18.5 && bmi <= 22.9:
		return "Berat badan normal"
	default:
		return "Kelebihan berat badan"
	}
}

func main() {
	var berat, tinggi int
	fmt.Print("Masukkan tinggi badan (cm): ")
	fmt.Scan(&tinggi)
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&berat)
	
	hasil := hitungBMI(berat, tinggi)
	fmt.Printf("Kategori: %s\n", hasil)
}
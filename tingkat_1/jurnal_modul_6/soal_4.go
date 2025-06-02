package main

import (
	"fmt"
)

func persamaanValid(koefA, koefB, koefC float64) bool {
	if koefA == 0 {
		return false
	}

	diskriminan := koefB*koefB - 4*koefA*koefC

	return diskriminan > 0
}

func main4() {
	var jumlahPersamaan int
	fmt.Print("Masukkan jumlah persamaan N: ")
	fmt.Scan(&jumlahPersamaan)

	hasil := make([]bool, jumlahPersamaan)

	for i := 0; i < jumlahPersamaan; i++ {
		var koefA, koefB, koefC float64
		fmt.Printf("Masukkan koefisien a, b, c untuk persamaan ke-%d: ", i+1)
		fmt.Scan(&koefA, &koefB, &koefC)
		hasil[i] = persamaanValid(koefA, koefB, koefC)
	}

	fmt.Println("Keluaran:")
	for _, valid := range hasil {
		fmt.Println(valid)
	}
}

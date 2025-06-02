package main

import (
	"fmt"
	"math"
)

func main3() {
	var jumlahEsKrim, karungPer1000Gram, jumlahKarung int

	fmt.Println("Masukkan jumlah gram es krim yang diinginkan (positif):")
	fmt.Scan(&jumlahEsKrim)
	fmt.Println("Masukkan nilai x (karung susu per 1000 gram):")
	fmt.Scan(&karungPer1000Gram)

	jumlahKarung = int(math.Ceil(float64(jumlahEsKrim) / 1000.0 * float64(karungPer1000Gram)))

	fmt.Printf("Jumlah karung susu yang dibutuhkan: %d karung\n", jumlahKarung)
}

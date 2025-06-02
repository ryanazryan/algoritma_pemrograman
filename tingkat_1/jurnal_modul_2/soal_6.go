package main

import (
	"fmt"
	"math"
)

func main6() {
	var rAlas, tinggiPiala float64

	fmt.Print("Masukkan jari-jari alas piala: ")
	fmt.Scan(&rAlas)
	fmt.Print("Masukkan tinggi piala: ")
	fmt.Scan(&tinggiPiala)

	rBola := (2.0 / 3.0) * rAlas

	tinggikerucut := tinggiPiala - rBola*2
	volumeKerucut := (1.0 / 3.0 ) * math.Pi * math.Pow(rAlas, 2) * tinggikerucut

	volumeBola := (4.0 / 3.0) * math.Pi * math.Pow(rBola, 3)

	volumePiala := volumeKerucut + volumeBola

	fmt.Printf("Volume piala adalah %.13f\n", volumePiala)

	// Penjelasan singkat, maaf untuk output tidak sangat presisi / persis sekali dengan example output yang diberikan
	// volume kerucut = 1/3 x phi x r^2 alas x tinggi kerucut
	// volume bola = 4/3 x phi x r^3 bola
}

package main

import (
	"fmt"
	"math"
)

func hitungFungsi(x, y, alpha float64) float64 {
	return (1-alpha)*math.Pow(x, (3/(2*y))) + alpha*math.Pow((4*y)/(5*x), 1)
}

func main() {
	var x1, y1, alpha1 float64
	var x2, y2, alpha2 float64

	fmt.Println("Masukkan nilai x1, y1, dan alpha1 (pisahkan dengan spasi):")
	fmt.Scan(&x1, &y1, &alpha1)

	fmt.Println("Masukkan nilai x2, y2, dan alpha2 (pisahkan dengan spasi):")
	fmt.Scan(&x2, &y2, &alpha2)

	hasil1 := hitungFungsi(x1, y1, alpha1)
	hasil2 := hitungFungsi(x2, y2, alpha2)

	fmt.Printf("Hasil untuk baris pertama: %.5f\n", hasil1)
	fmt.Printf("Hasil untuk baris kedua: %.5f\n", hasil2)
}

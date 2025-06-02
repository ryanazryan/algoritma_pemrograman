package main

import (
	"fmt"
)

func hitungFungsi(x, y, alpha float64) float64 {
	return (1-alpha)*(3*x/(2*y)) + alpha*(4*y/(5*x))
}

func main() {
	var x1, y1, alpha1 float64
	var x2, y2, alpha2 float64

	fmt.Print("Masukkan nilai x1, y1, dan alpha1: ")
	fmt.Scan(&x1, &y1, &alpha1)

	fmt.Print("Masukkan nilai x2, y2, dan alpha2: ")
	fmt.Scan(&x2, &y2, &alpha2)

	hasil1 := hitungFungsi(x1, y1, alpha1)
	hasil2 := hitungFungsi(x2, y2, alpha2)

	fmt.Printf("Hasil:\n%.5f\n%.5f\n", hasil1, hasil2)
}

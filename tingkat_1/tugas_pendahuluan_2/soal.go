package main

import "fmt"

func menghitung(x,y float64)float64{
	return 5/(2*y+7) + x*x - 3*y
}

func main7(){
	var n1, n2 float64
	fmt.Print("Masukkan nilai n1: ")
	fmt.Scan(&n1)
	fmt.Print("Masukkan nilai n2: ")
	fmt.Scan(&n2)

	hasil1 := menghitung(n1, n2)
	hasil2 := menghitung(n2, n1)

	fmt.Printf("Hasil f(n1, n2): %.6f\n", hasil1)
	fmt.Printf("Hasil f(n2, n1): %.6f\n", hasil2)
}
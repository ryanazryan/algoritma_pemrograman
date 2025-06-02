package main

import "fmt"

func main1(){
	type persegi struct {
		sisi float64
		keliling float64
		luas float64
	}

	var jumlah int
	fmt.Scan(&jumlah)

	var square persegi
	for i := 1; i <= jumlah; i++ {
		fmt.Scan(&square.sisi)

		square.keliling = 4 * square.sisi
		square.luas = square.sisi * square.sisi

		fmt.Println(square.keliling, square.luas)
	}
}
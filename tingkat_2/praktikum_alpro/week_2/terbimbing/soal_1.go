package main

import "fmt"

func main1(){
	type persegiPanjang struct{
		panjang float64
		lebar float64
		keliling float64
		luas float64
	}

	var jumlah, i int
	fmt.Scan(&jumlah)

	var rectangle persegiPanjang
	for i = 1; i <= jumlah; i++ {
		fmt.Scan(&rectangle.panjang)
		fmt.Scan(&rectangle.lebar)
		rectangle.keliling = (2*rectangle.panjang) + (2*rectangle.lebar)
		rectangle.luas = rectangle.panjang*rectangle.lebar
		fmt.Println(rectangle.keliling, rectangle.luas)
	}
}
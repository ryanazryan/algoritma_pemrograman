package main

import "fmt"

func main() {
	var n, panjang, lebar, luas int

	fmt.Println("Masukkan jumlah persegi panjang:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan panjang dan lebar untuk persegi panjang ke-%d: ", i)
		fmt.Scan(&panjang, &lebar)

		luas = panjang * lebar
		if luas < 1000 {
			fmt.Println(luas)
		} else {
			break
		}
	}
}

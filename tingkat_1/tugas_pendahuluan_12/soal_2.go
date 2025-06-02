package main

import (
	"fmt"
)

func main2() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	count := 0
	for angka > 0 {
		angka /= 10
		count++
	}

	fmt.Printf("%d digit\n", count)
}
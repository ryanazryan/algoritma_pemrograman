package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan (1-999999): ")
	fmt.Scan(&number)

	switch {
	case number >= 1 && number <= 9:
		fmt.Println("Satuan")
	case number >= 10 && number <= 99:
		fmt.Println("Puluhan")
	case number >= 100 && number <= 999:
		fmt.Println("Ratusan")
	case number >= 1000 && number <= 9999:
		fmt.Println("Ribuan")
	case number >= 10000 && number <= 99999:
		fmt.Println("Puluhan ribu")
	case number >= 100000 && number <= 999999:
		fmt.Println("Ratusan ribu")
	default:
		fmt.Println("Bilangan tidak valid")
	}
}

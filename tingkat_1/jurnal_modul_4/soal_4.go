package main

import (
	"fmt"
	"strconv"
)

func jumlahDigit(bilangan int) int {
	bilStr := strconv.Itoa(bilangan)
	total := 0

	for _, digit := range bilStr {
		angka, _ := strconv.Atoi(string(digit))
		total += angka
	}

	return total
}

func main4() {
	bilanganList := []int{12345, 98142, 99999}

	for _, bilangan := range bilanganList {
		fmt.Printf("Bilangan: %d, Hasil penjumlahan nilai setiap digit adalah %d\n", bilangan, jumlahDigit(bilangan))
	}
}
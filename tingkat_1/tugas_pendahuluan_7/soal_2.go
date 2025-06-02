package main

import (
	"fmt"
	"strconv"
)

func main2() {
	var angka int
	fmt.Print("Masukkan bilangan tiga digit: ")
	fmt.Scan(&angka)

	hasil := gandakanDigit(angka)
	fmt.Println("Hasil:", hasil)
}

func gandakanDigit(angka int) int {
	strAngka := strconv.Itoa(angka)
	var hasilStr string

	for _, digit := range strAngka {
		hasilStr += string(digit) + string(digit)
	}

	hasil, _ := strconv.Atoi(hasilStr)
	return hasil
}

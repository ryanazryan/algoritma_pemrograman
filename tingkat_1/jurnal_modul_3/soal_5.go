package main

import (
	"fmt"
)

func cekMenang(bil1, bil2, bil3 int) bool {
	jumlahGenap := (bil1+bil3)%2 == 0
	keduaGanjil := bil2%2 != 0

	return jumlahGenap && keduaGanjil
}

func main5() {
	var bil1, bil2, bil3 int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&bil1)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&bil2)
	fmt.Print("Masukkan bilangan ketiga: ")
	fmt.Scan(&bil3)

	if cekMenang(bil1, bil2, bil3) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
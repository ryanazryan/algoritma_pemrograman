package main

import (
	"fmt"
	"strconv"
)

func hitungJumlahGanjilGenap(angka int) (int, int) {
    if angka == 0 {
        return 0, 0
    }

    jumlahGanjil := 0
    jumlahGenap := 0

    angkaStr := strconv.Itoa(abs(angka))

    for _, digit := range angkaStr {
        digitInt, _ := strconv.Atoi(string(digit))
        if digitInt%2 == 0 {
            jumlahGenap++
        } else {
            jumlahGanjil++
        }
    }

    return jumlahGanjil, jumlahGenap
}


func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var masukan int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scanln(&masukan)

	jumlahGanjil, jumlahGenap := hitungJumlahGanjilGenap(masukan)
	fmt.Printf("%d %d\n", jumlahGanjil, jumlahGenap)
}

package main

import (
	"fmt"
	"math"
)

func main1() {
	// Input bilangan tebakan adik dan kakak
	var adik, kakak int
	fmt.Print("Masukkan bilangan tebakan adik: ")
	fmt.Scan(&adik)
	fmt.Print("Masukkan bilangan tebakan kakak: ")
	fmt.Scan(&kakak)

	// Panggil fungsi untuk menentukan pemenang
	menang := cekPemenang(adik, kakak)
	fmt.Println(menang)
}

// Fungsi cekPemenang untuk menentukan apakah adik menang
func cekPemenang(adik, kakak int) bool {
	// Adik menang jika tebakan sama atau selisihnya 1 atau 5
	return adik == kakak ||
		math.Abs(float64(adik-kakak)) == 1 ||
		math.Abs(float64(adik-kakak)) == 5
}

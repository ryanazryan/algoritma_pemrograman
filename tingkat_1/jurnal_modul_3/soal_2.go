package main

import (
	"fmt"
)

func bandingkanNilai(nilaiMataUangAsing float64, namaMataUang string, nilaiSatuanRupiah, nilaiRupiahDiberikan float64) string {
	nilaiDalamRupiah := nilaiMataUangAsing * nilaiSatuanRupiah

	if nilaiDalamRupiah > nilaiRupiahDiberikan {
		return fmt.Sprintf("%g %s mahal? true", nilaiMataUangAsing, namaMataUang)
	}
	return fmt.Sprintf("%g %s mahal? false", nilaiMataUangAsing, namaMataUang)
}

func main2() {
	var nilaiMataUangAsing, nilaiSatuanRupiah, nilaiRupiahDiberikan float64
	var namaMataUang string

	fmt.Print("Masukkan nilai mata uang asing: ")
	fmt.Scan(&nilaiMataUangAsing)
	fmt.Print("Masukkan nama mata uang (contoh: USD, CNY, JPY): ")
	fmt.Scan(&namaMataUang)
	fmt.Print("Masukkan nilai kurs dalam rupiah: ")
	fmt.Scan(&nilaiSatuanRupiah)
	fmt.Print("Masukkan nilai rupiah yang akan dibandingkan: ")
	fmt.Scan(&nilaiRupiahDiberikan)

	fmt.Println(bandingkanNilai(nilaiMataUangAsing, namaMataUang, nilaiSatuanRupiah, nilaiRupiahDiberikan))
}
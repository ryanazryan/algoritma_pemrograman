package main

import (
	"fmt"
	"strings"
)

func rataRataPanjangEmpatKata(kalimat string) float64 {
	kataList := strings.Fields(kalimat)

	if len(kataList) != 4 {
		fmt.Println("Kalimat harus terdiri dari tepat 4 kata.")
		return 0.0
	}

	totalPanjang := 0
	for _, kata := range kataList {
		totalPanjang += len(kata)
	}

	return float64(totalPanjang) / 4
}

func main5() {
	kalimatList := []string{
		"Saya anak rajin membaca",           
		"Hari ini cuaca mendung",            
		"Naufal Faiq Azryan",    // karena nama lengkap saya hanya 3 kata
	}

	for _, kalimat := range kalimatList {
		hasil := rataRataPanjangEmpatKata(kalimat)
		fmt.Printf("Kalimat: \"%s\", Rata-rata panjang kata: %.2f\n", kalimat, hasil)
	}
}
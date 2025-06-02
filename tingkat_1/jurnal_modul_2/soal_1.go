package main

import "fmt"

func main1(){
	var panjang, lebar, panjangAlas, tinggiSegitiga, tinggiBidang float64

	fmt.Print("Masukkan panjang dan lebar persegi panjang : ")
	fmt.Scan(&panjang, &lebar)

	fmt.Print("Masukkan panjang alas dan tinggi segitiga: ")
	fmt.Scan(&panjangAlas, &tinggiSegitiga)

	fmt.Print("Masukkan tinggi bidang 3D: ")
	fmt.Scan(&tinggiBidang)

	luasPersegiPanjang := panjang * lebar
	luasSegitiga := 0.5 * panjangAlas * tinggiSegitiga

	totalLuas := luasPersegiPanjang + luasSegitiga

	volume := totalLuas * tinggiBidang

	fmt.Printf("Volume dari bidang 3D adalah : %.2f\n", volume)

	// Penjelasan algoritma singkat

	// rumus menghitung volume dari bentuk gabungan
	// rumus menghitung volume luas persegi panjang: panjang x lebar = 10 x 10 = 100
	// luas segitiga = 1/2 x panjang alas x tinggi: 1/2 x 3 x 5 = 7.5
	// total luas alas = luas persegi panjang + luas segitiga: 100 + 7.5 = 107.5

	// setelah mendapatkan total luas, hitung volume bentuk 3D dengan mengalikan total luas dengan tinggi keseluruhan
	// volume = total luas x tinggi bidang = 107.5 = 537.5
	// catatan : di contoh masukan dan keluaran	output pada bagian 2 547.5 , hasil perhitungan yang tepat adalah 537.5
}
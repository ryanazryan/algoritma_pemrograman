package main

import "fmt"

func main3(){
	var hargaSabun, hargaSampo, hargaPastaGigi float64

	fmt.Print("Masukkan harga sabun: ")
	fmt.Scan(&hargaSabun)
	fmt.Print("Masukkan harga sampo: ")
	fmt.Scan(&hargaSampo)
	fmt.Print("Masukkan harga pasta gigi : ")
	fmt.Scan(&hargaPastaGigi)

	hargaSabunSetelahDiskon := hargaSabun - (hargaSabun * 0.05)
	hargaSampoSetelahDiskon := hargaSampo - (hargaSampo * 0.10)
	hargaPastaGigiSetelahDiskon := hargaPastaGigi- (hargaPastaGigi * 0.08)

	totalBayar := hargaSabunSetelahDiskon + hargaSampoSetelahDiskon + hargaPastaGigiSetelahDiskon

	fmt.Printf("Total belanja setelah diskon: %d\n", int(totalBayar))

	// Penjelasan singkat
	// harga setelah diskon = harga asli - ( harga asli x diskon)
	// setelah menghitung harga masing" barang setelah diskon, hitung total biaya
	// total biaya = harga sabun setelah diskon + harga sampo setelah diskon + harga pasta gigi setelah diskon
}
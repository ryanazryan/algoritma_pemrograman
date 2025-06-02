package main

import (
	"fmt"
)

var jenisKendaraan string
var jam1, menit1, detik1 int
var jam2, menit2, detik2 int 
var totalUang int

func menu() {
	fmt.Println("----------------------------")
	fmt.Println("           M E N U           ")
	fmt.Println("----------------------------")
	fmt.Println("1. Input Kendaraan Masuk")
	fmt.Println("2. Input Kendaraan Keluar")
	fmt.Println("3. Hitung Biaya Parkir")
	fmt.Println("4. Cetak Total Uang")
	fmt.Println("5. Exit")
	fmt.Println("----------------------------")
}

func inputKendaraanMasuk() {
	fmt.Print("Masukkan jenis kendaraan (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Masukkan jam, menit, detik kendaraan masuk: ")
	fmt.Scan(&jam1, &menit1, &detik1)
}

func inputKendaraanKeluar() {
	fmt.Print("Masukkan jam, menit, detik kendaraan keluar: ")
	fmt.Scan(&jam2, &menit2, &detik2)
}

func hitungBiayaParkir() {
	durasiJam := jam2 - jam1
	durasiMenit := menit2 - menit1
	durasiDetik := detik2 - detik1

	if durasiDetik < 0 {
		durasiDetik += 60
		durasiMenit--
	}
	if durasiMenit < 0 {
		durasiMenit += 60
		durasiJam--
	}

	if durasiJam < 0 {
		durasiJam = 0
	}

	var biayaParkir int
	if jenisKendaraan == "mobil" {
		biayaParkir = 5000 + (durasiJam-1)*3000
	} else if jenisKendaraan == "motor" {
		biayaParkir = 2000 + (durasiJam-1)*1000
	} else {
		fmt.Println("Jenis kendaraan tidak dikenali")
		return
	}

	fmt.Printf("Biaya parkir %s selama %d jam: Rp %d,-\n", jenisKendaraan, durasiJam, biayaParkir)
	totalUang += biayaParkir
}

func cetakTotalUang() {
	fmt.Printf("Total uang: Rp %d,-\n", totalUang)
}

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			inputKendaraanMasuk()
		case 2:
			inputKendaraanKeluar()
		case 3:
			hitungBiayaParkir()
		case 4:
			cetakTotalUang()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih antara 1-5.")
		}
	}
}
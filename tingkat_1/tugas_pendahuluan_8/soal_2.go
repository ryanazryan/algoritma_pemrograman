package main

import (
	"fmt"
)

func main2() {
	var memilikiKartuIdentitas, memilikiSuratIzin bool

	fmt.Println("Apakah pengunjung memiliki Kartu Identitas? (true/false):")
	fmt.Scan(&memilikiKartuIdentitas)
	fmt.Println("Apakah pengunjung memiliki Surat Izin? (true/false):")
	fmt.Scan(&memilikiSuratIzin)

	if memilikiKartuIdentitas {
		fmt.Println("Pengunjung diizinkan masuk. Kartu Identitas ada.")
	} else if memilikiSuratIzin {
		fmt.Println("Pengunjung diizinkan masuk. Surat Izin ada.")
	} else {
		fmt.Println("Pengunjung tidak diizinkan masuk. Kartu Identitas dan Surat Izin tidak ada.")
	}
}

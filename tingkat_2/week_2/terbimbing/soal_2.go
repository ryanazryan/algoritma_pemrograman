package main

import "fmt"

func main() {
	type creator struct {
		nama     string
		listener int
		song     int	
	}

	var c creator
	fmt.Scan(&c.nama, &c.listener, &c.song)

	var selisihLagu, selisihListener int
	selisihLagu = 10 - c.song
	selisihListener = 500000 - c.listener

	if selisihLagu <= 0 && selisihListener <= 0 {
		fmt.Println("Kreator", c.nama, "dapat terverifikasi dengan")
		fmt.Println("jumlah total lagu: ", c.song)
		fmt.Println("pendengar bulanan sebanyak: ", c.listener)
	} else if selisihLagu > 0 && selisihListener > 0 {
		fmt.Println("Kreator", c.nama, "belum dapat terverifikasi butuh")
		fmt.Println("jumlah total lagu: ", selisihLagu)
		fmt.Println("pendengar bulanan sebanyak: ", selisihListener, "lagi")
	} else if selisihLagu <= 0 && selisihListener > 0 {
		fmt.Println("Kreator", c.nama, "Belum dapat terverfirikasi, butuh", selisihListener, "pendengar bulanan lagi")
	} else {
		fmt.Println("Kreator", c.nama, "belum dapat terverifikasi, butuh, ", selisihLagu, "Lagu rilis lagi")
	}
}

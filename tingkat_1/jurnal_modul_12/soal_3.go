package main

import "fmt"

func main3() {
	var skor, percobaan int
	percobaan = 1

	for percobaan <= 5 {
		fmt.Printf("Masukkan skor EPrT ke-%d: ", percobaan)
		fmt.Scan(&skor)

		fmt.Printf("Skor EPrT ke-%d Anda %d\n", percobaan, skor)
		if skor >= 500 {
			fmt.Println("Selamat Anda bisa mendaftar sidang Yudisium")
			break
		}

		percobaan++
	}

	if percobaan > 5 && skor < 500 {
		fmt.Println("Anda telah mencapai batas percobaan.")
	}
}

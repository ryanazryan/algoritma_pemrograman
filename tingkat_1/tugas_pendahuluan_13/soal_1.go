package main

import (
	"fmt"
)

func main1() {
	const PIN_BENAR = 123987
	kesempatan := 3

	for kesempatan > 0 {
		var PIN int
		fmt.Print("Masukkan PIN: ")
		fmt.Scan(&PIN)

		if PIN == PIN_BENAR {
			var nominal int
			fmt.Print("Masukkan nominal uang: ")
			fmt.Scan(&nominal)
			fmt.Printf("Uang Rp%d sudah bisa diambil, terima kasih.\n", nominal)
			break
		} else {
			kesempatan--
			if kesempatan > 0 {
				fmt.Println("PIN salah")
			} else {
				fmt.Println("PIN terblokir, silahkan kunjungi kantor cabang terdekat!")
			}
		}
	}
}

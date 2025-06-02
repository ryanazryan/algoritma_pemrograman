package main

import (
	"fmt"
)

func main1() {
	var username, password string
	var jumlahGagal int

	for {
		fmt.Println("Masukkan username dan password:")
		fmt.Scan(&username, &password)

		if username == "admin" && password == "admin" {
			break
		}
		jumlahGagal++
	}

	fmt.Println(jumlahGagal)
	fmt.Println("Login berhasil")
}

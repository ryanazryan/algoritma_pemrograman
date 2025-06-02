package main

import (
	"fmt"
	"strings"
)

func main1() {
	var warna1, warna2, temp string
	fmt.Println("Masukkan dua warna primer yang akan dicampur (merah, kuning, biru):")
	fmt.Scan(&warna1, &warna2)

	warna1 = strings.ToLower(warna1)
	warna2 = strings.ToLower(warna2)

	if warna1 > warna2 {
		temp = warna1
		warna1 = warna2
		warna2 = temp
	}

	if warna1 == warna2 {
		fmt.Printf("Hasil: %s\n", warna1)
	} else {
		switch warna1 {
		case "biru":
			switch warna2 {
			case "kuning":
				fmt.Println("Hasil: hijau")
			case "merah":
				fmt.Println("Hasil: ungu")
			default:
				fmt.Println("Warna tidak valid.")
			}
		case "kuning":
			switch warna2 {
			case "merah":
				fmt.Println("Hasil: jingga")
			default:
				fmt.Println("Warna tidak valid.")
			}
		case "merah":
			switch warna2 {
			case "biru":
				fmt.Println("Hasil: ungu")
			case "kuning":
				fmt.Println("Hasil: jingga")
			default:
				fmt.Println("Warna tidak valid.")
			}
		default:
			fmt.Println("Warna tidak valid.")
		}
	}
}

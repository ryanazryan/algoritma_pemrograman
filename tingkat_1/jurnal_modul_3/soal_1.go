package main

import (
	"fmt"
)

func cekFaktor(bilangan1, bilangan2 int) string {
	if bilangan1%bilangan2 == 0 || bilangan2%bilangan1 == 0 {
		return fmt.Sprintf("Apakah %d dan %d adalah faktor? true", bilangan1, bilangan2)
	}
	return fmt.Sprintf("Apakah %d dan %d adalah faktor? false", bilangan1, bilangan2)
}

func main1() {
	var bilangan1, bilangan2 int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&bilangan1)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&bilangan2)

	fmt.Println(cekFaktor(bilangan1, bilangan2))
}
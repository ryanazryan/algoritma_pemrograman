package main

import (
	"fmt"
	"strings"
)

func main1() {
	var value float64
	var unit string

	fmt.Print("Masukkan suhu (contoh: 30 Celcius atau 86 Fahrenheit): ")
	_, err := fmt.Scanf("%f %s", &value, &unit)
	if err != nil {
		fmt.Println("Input tidak valid. Pastikan format input benar.")
		return
	}

	unit = strings.ToLower(unit)

	if unit == "celcius" {
		fahrenheit := (value * 9 / 5) + 32
		fmt.Printf("Suhu dalam Fahrenheit adalah %.2f F\n", fahrenheit)
	} else if unit == "fahrenheit" {
		celcius := (value - 32) * 5 / 9
		fmt.Printf("Suhu dalam Celcius adalah %.2f C\n", celcius)
	} else {
		fmt.Println("Satuan suhu tidak valid. Gunakan 'Celcius' atau 'Fahrenheit'.")
	}
}

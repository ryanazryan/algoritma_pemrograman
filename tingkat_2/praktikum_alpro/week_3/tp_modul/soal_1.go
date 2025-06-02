package main

import (
	"fmt"
)

func keReamur(celsius float64) float64 {
	return 4.0 / 5.0 * celsius
}

func keFahrenheit(celsius float64) float64 {
	return 9.0/5.0*celsius + 32
}

func keKelvin(celsius float64) float64 {
	return celsius + 273
}

func main1() {
	var awal, akhir, langkah float64

	fmt.Scan(&awal, &akhir, &langkah)

	fmt.Println("Celcius    Reamur    Fahrenheit    Kelvin")

	for suhu := awal; suhu <= akhir; suhu += langkah {
		reamur := keReamur(suhu)
		fahrenheit := keFahrenheit(suhu)
		kelvin := keKelvin(suhu)

		fmt.Printf("%-10.2f %-8.2f %-12.2f %.2f\n", suhu, reamur, fahrenheit, kelvin)
	}
}

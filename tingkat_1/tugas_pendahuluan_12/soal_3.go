package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	adalahKonsekutif := true
	sebelumnya := bilangan % 10
	bilangan /= 10

	for bilangan > 0 {
		sekarang := bilangan % 10
		if int(math.Abs(float64(sekarang-sebelumnya))) != 1 {
			adalahKonsekutif = false
			break
		}
		sebelumnya = sekarang
		bilangan /= 10
	}

	fmt.Println(adalahKonsekutif)
}

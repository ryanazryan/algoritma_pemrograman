package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Masukkan dua angka: ")
	fmt.Scan(&a, &b)

	min, max := a, b
	if a > b {
		min, max = b, a
	}

	total := 0
	for i := min; i <= max; i++ { fmt.Print(i, " ")
		total += i
	}

	fmt.Printf("\nTotal = %d\n", total)
}

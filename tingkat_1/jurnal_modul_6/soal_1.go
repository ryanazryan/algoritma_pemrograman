package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	y := 1.0

	iterations := 10

	for i := 0; i < iterations; i++ {
		y = (y + x/y) / 2
	}

	return math.Round(y*1e8) / 1e8
}

func main1() {
	inputs := []int{2, 3, 4, 5, 64}

	for _, x := range inputs {
		fmt.Printf("Masukan: %d, Keluaran: %.8f\n", x, sqrt(float64(x)))
	}
}

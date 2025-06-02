package main

import (
	"fmt"
)

func isYinDominan(sequence string) bool {
	jumlahYin := 0
	jumlahYang := 0

	for _, char := range sequence {
		if char == '1' {
			jumlahYin++
		} else if char == '0' {
			jumlahYang++
		}
	}

	return jumlahYin > jumlahYang
}

func main2() {
	sekuens := []string{
		"11100101101", 
		"11010011000",
		"10010011010",
		"11111100000", 
		"00000011110",
	}

	for _, seq := range sekuens {
		hasil := isYinDominan(seq)
		fmt.Printf("Sekuens: %s, Apakah Yin lebih dominan? %v\n", seq, hasil)
	}
}
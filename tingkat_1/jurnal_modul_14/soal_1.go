package main

import (
	"fmt"
)

func main1() {
	var n, i, j int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	i = 1
	for i <= n {
		for j = 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
		i = i + 1
	}
}

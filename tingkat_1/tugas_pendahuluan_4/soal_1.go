package main

import (
	"fmt"
)

func main() {
	var x, N int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	if x%N == 0 {
		fmt.Printf("%d kelipatan %d? true\n", x, N)
	} else {
		fmt.Printf("%d kelipatan %d? false\n", x, N)
	}
}

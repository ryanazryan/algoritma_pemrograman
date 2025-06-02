package main

import "fmt"

func cetakBaris(n int) {
	if n%2 != 0 {
		n = n - 1
	}

	fmt.Print(n)

	if n > 2 {
		fmt.Print(", ")
		cetakBaris(n - 2)
	}
}

func main2(){
	var n int
	fmt.Scan(&n)
	cetakBaris(n)
	fmt.Println()
}
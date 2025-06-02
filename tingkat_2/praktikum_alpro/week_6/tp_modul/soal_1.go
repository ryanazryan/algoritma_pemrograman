package main

import "fmt"

func fibolagifibolagi(n int) int {
	if n == 1{
		return 6
	} else if n == 2{
		return 7
	} else if n > 2 {
		return fibolagifibolagi(n-1) + fibolagifibolagi(n-2) + fibolagifibolagi(n-3)
	}

	return 0
}

func main1() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibolagifibolagi(n))
}
package main

import "fmt"

func f(n int) int {
	if n == 0 {
		return 1
	}

	return f(n - 1) *n
}

func main1(){
	var n int
	fmt.Scan(&n)

	fmt.Print(f(n))
}
package main

import "fmt"

func pangkat(a, b int) int {
	if b == 1 {
		return a
	}

	return a * pangkat(a, b-1)
}

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(pangkat(a, b))
}
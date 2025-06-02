package main

import (
	"fmt"
	"strconv"
)

func main2() {
	var n int
	fmt.Println("Masukkan bilangan bulat positif N:")
	fmt.Scan(&n)

	result := ""
	for i := n; i > 0; i-- {
		result += strconv.Itoa(i)
		if i > 1 {
			result += " x "
		}
	}

	fmt.Println(result)
}

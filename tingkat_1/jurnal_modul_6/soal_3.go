package main

import (
	"fmt"
	"strings"
)

func main3() {
	var x int
	var p string

	fmt.Print("Masukkan digit x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan 10 digit p: ")
	fmt.Scan(&p)

	xStr := fmt.Sprintf("%d", x)

	contains := strings.Contains(p, xStr)

	fmt.Println(contains)
}

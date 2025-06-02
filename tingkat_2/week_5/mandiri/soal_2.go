package main

import "fmt"

func E(n int) int {
	if n == 2{
		return 2
	}
	return E(n-2) + n
}

func main2(){
	var num int
	fmt.Scan(&num)
	fmt.Println(E(num))
}
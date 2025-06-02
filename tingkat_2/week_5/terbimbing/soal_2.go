package main

import "fmt"

func jumlahDigit(n int) int {
	if n == 0 {
		return 0
	}

	return (n%10) + jumlahDigit(n/10)
}

func main2(){
	var num int
	fmt.Scan(&num)
	fmt.Println(jumlahDigit(num))
}
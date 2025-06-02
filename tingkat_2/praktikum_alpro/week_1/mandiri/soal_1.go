package main

import "fmt"

func main1(){
	var X int
	var oddCount, evenCount int

	fmt.Scan(&X)

	for X > 0 {
		digit := X % 10
		if digit % 2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		X /= 10
	}

	fmt.Println(oddCount,evenCount)
}
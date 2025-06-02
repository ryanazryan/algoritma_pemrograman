package main
import "fmt"

func fibofiboan(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n == 3{
		return 4
	} else {
		return fibofiboan(n-1) + fibofiboan(n-2) + fibofiboan(n-3)
	}
}

func main1(){
	var n int
	fmt.Scan(&n)
	fmt.Println(fibofiboan(n))
}
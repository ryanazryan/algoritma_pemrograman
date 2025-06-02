package main
import "fmt"

func lucas(n int) int {
	if n == 1{
		return 2
	} else if n == 2{
		return 1
	} else {
		return lucas(n-1) + lucas(n-2)
	}
}

func main1(){
	var n int
	fmt.Scan(&n)
	fmt.Println(lucas(n))
}
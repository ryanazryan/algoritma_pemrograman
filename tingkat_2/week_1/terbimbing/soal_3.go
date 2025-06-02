package main

import "fmt"

func main(){
	var N, a, b,temp, i int
	
	fmt.Scan(&N)

	a = 1
	b = 1
	i = 0

	for i < N {
		if(a%2!=0){
			fmt.Print(a, "")
			i++
		}
		temp = a + b
		a = b
		b = temp
	}
}
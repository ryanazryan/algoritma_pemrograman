package main

import "fmt"

func main2(){
	var X, i int
	var isPrime bool
	
	fmt.Scan(&X)
	
	isPrime = X > 1

	if isPrime == true {
		for i = 2; i < X; i++ {
		 isPrime = (isPrime == true) &&  (X%i != 0)
		}
	}
	
	if isPrime == true {
		fmt.Println(X, "adalah bilangan prima")
	} else {
		fmt.Println(X, "Bukan bilangan prima")
	}
}
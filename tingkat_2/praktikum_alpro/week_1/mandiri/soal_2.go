package main

import "fmt"

func main(){
	var N, reversed int
	var isPalindrome bool

	fmt.Scan(&N)

	reversed = 0
	original := N

	for N > 0 {
		digit := N % 10
		reversed = reversed * 10 + digit
		N /= 10
	}

	isPalindrome = (original) == reversed

	if isPalindrome {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}
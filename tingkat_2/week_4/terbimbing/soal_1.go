package main

import (
	"fmt"
)

func tampilMenu() {
	fmt.Println("-----------------------------")
	fmt.Println("          M E N U            ")
	fmt.Println("-----------------------------")
	fmt.Println("1. Hello")
	fmt.Println("2. Hi")
	fmt.Println("3. Good Morning")
	fmt.Println("4. Good Night")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------------")
}

func main1() {
	var pilihan int

	for {
		tampilMenu()
		fmt.Print("Your Choice (1/2/3/4/5)? ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Hello")
		case 2:
			fmt.Println("Hi")
		case 3:
			fmt.Println("Good Morning")
		case 4:
			fmt.Println("Good Night")
		case 5:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Invalid choice. Please choose between 1-5.")
		}
	}
}
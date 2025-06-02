package main

import (
	"fmt"
	"math/rand"
	"time"
)

func yourGuessing(yN *int) {
	fmt.Print("Enter your guess: ")
	fmt.Scan(yN)
}

func numGenerator(cN *int) {
	rand.Seed(time.Now().UnixNano())
	*cN = rand.Intn(5)
}

func process(yN, cN int, win *bool) {
	*win = false
	if yN == cN {
		*win = true
	}
	fmt.Printf("Your guessing: %d, computer number: %d, win: %v\n", yN, cN, *win)
}

func conclusion(r int, win bool) {
	if win {
		fmt.Printf("You win in %d round\n", r)
	} else {
		fmt.Printf("Computer win in %d round\n", r)
	}
}

func main() {
	var round int
	var computerNumber, yourNumber int
	var win bool

	round = 0
	win = false
	fmt.Println("Start")

	for round < 5 && !win {
		round++
		fmt.Printf("Round %d\n", round)

		numGenerator(&computerNumber)

		yourGuessing(&yourNumber)

		process(yourNumber, computerNumber, &win)
	}

	conclusion(round, win)

	fmt.Println("End")
}
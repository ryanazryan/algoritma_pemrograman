package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main1() {
	scanner := bufio.NewScanner(os.Stdin)
	var prev, current int
	isAscending := true

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif atau 9999):")

	scanner.Scan()
	firstInput, _ := strconv.Atoi(scanner.Text())
	if firstInput < 0 || firstInput == 9999 {
		fmt.Println(true)
		return
	}
	prev = firstInput

	for {
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())
		if input < 0 || input == 9999 {
			break
		}

		current = input
		if current < prev {
			isAscending = false
		}
		prev = current
	}

	fmt.Println(isAscending)
}

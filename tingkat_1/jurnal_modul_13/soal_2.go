package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main2() {
	scanner := bufio.NewScanner(os.Stdin)
	var oddCount, evenCount int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")

	for {
		scanner.Scan()
		input := scanner.Text()
		number, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		if number < 0 {	
			break
		}

		if number%2 == 0 {
			fmt.Printf("%d genap\n", number)
			evenCount++
		} else {
			fmt.Printf("%d ganjil\n", number)
			oddCount++
		}
	}

	fmt.Printf("%d genap, %d ganjil\n", evenCount, oddCount)
}

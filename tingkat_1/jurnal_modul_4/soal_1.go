package main

import (
	"fmt"
	"strconv"
)

func isAscending(number int) bool {
	numStr := strconv.Itoa(number)

	if len(numStr) != 4 {
		return false
	}

	count := 1
	for i := 0; i < len(numStr)-1; i++ {
		d1 := int(numStr[i] - '0')
		d2 := int(numStr[i+1] - '0')

		if d1 < d2 {
			count++
			if count >= 3 {
				return true
			}
		} else {
			count = 1
		}
	}

	return false
}

func main1() {
	numbers := []int{2345, 3456, 4321, 9876, 2468, 1357, 6423, 1232}

	for _, num := range numbers {
		fmt.Printf("Number: %d, Ascending: %v\n", num, isAscending(num))
	}
}
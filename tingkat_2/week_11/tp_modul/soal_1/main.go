package main

func binSearch(tab []int, n int, x int) bool {
	left := 0
	right := n -1
	var mid int

	for left <= right {
		mid = (left + right) / 2

		if x < tab[mid] {
			right = mid -1
		} else if x > tab[mid] {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
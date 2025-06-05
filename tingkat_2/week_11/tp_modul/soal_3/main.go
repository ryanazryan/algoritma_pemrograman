package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func binarySearch(T tabInt, n int, x int) int {
	left, right := 0, n-1
	mid := -1

	for left <= right {
		mid = (left + right) / 2
		if x < T[mid] {
			right = mid - 1
		} else if x > T[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1 
}

func bacaData(A *tabInt, n *int) {
	fmt.Scan(*n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	fmt.Println("Data dalam array:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func main() {
	var data tabInt
	var nData int
	var x1 int

	fmt.Scan(&nData)
	bacaData(&data, &nData)
	fmt.Scan(&x1)
	cetakData(data, nData)
	index := binarySearch(data, nData, x1)
	if index != -1 {
		fmt.Printf("Data ditemukan pada indeks ke-%d\n", index)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
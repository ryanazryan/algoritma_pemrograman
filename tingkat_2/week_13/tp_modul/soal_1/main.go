package main

import "fmt"

const NMAX int = 99

type tabInt [NMAX]int

func bacaData(A *tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func insertionSort(A *tabInt, n int) {
	var temp int
	for pass := 1; pass < n; pass++ {
		i := pass
		temp = A[pass]
		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
	}
}

func main() {
	var data tabInt
	var nData int

	fmt.Scan(&nData)
	bacaData(&data, nData) 
	cetakData(data, nData) 
	insertionSort(&data, nData)
	cetakData(data, nData)    
}
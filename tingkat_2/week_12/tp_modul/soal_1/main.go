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

func selectionSort(A *tabInt, n int) {
	var i, idx, pass, temp int

	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if A[i] > A[idx] {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

func main() {
	var data tabInt
	var nData int

	fmt.Scan(&nData)
	bacaData(&data, nData)
	cetakData(data, nData)
	selectionSort(&data, nData)
	cetakData(data, nData)
}
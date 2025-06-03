package main

import "fmt"

const NMAX int = 1024

type arrPenjualan [NMAX]int

func main() {
	var data arrPenjualan
	var n, idxMax int

	for {
		var x int
		fmt.Scan(&x)
		if x == -1 {
			break
		}
		data[n] = x
		n++
	}

	idxMax = nilaiMax(data, n)
	fmt.Println(idxMax+1, data[idxMax])
}

func nilaiMax(A arrPenjualan, size int) int {
	idxMax := 0
	for i := 1; i < size; i++ {
		if A[i] > A[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}
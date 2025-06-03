package main

import "fmt"

const NMAX = 1024

func cariNilaiMax(arrBilangan [NMAX]int, nBilangan int) int {
	max := -9999999
	for i := 0; i < nBilangan; i++ {
		if arrBilangan[i] > max {
			max = arrBilangan[i]
		}
	}
	return max
}

func cariNilaiMin(arrBilangan [NMAX]int, nBilangan int) int {
	min := 9999999
	for i := 0; i < nBilangan; i++ {
		if arrBilangan[i] < min {
			min = arrBilangan[i]
		}
	}
	return min
}

func main() {
	var arrBilangan [NMAX]int
	var nBilangan int
	var bilangan int
	var i int = 0

	for {
		fmt.Scan(&bilangan)
		if bilangan == 0 {
			break
		}
		arrBilangan[i] = bilangan
		i++
		nBilangan++
	}

	max := cariNilaiMax(arrBilangan, nBilangan)
	min := cariNilaiMin(arrBilangan, nBilangan)

	fmt.Println(max, min)
}
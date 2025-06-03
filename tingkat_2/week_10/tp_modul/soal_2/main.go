package main

import "fmt"

const NMAX = 1024

type arrBaju [NMAX]string

func main() {
	var A arrBaju
	var searchKey string
	var size, i int
	var warnabaju string

	i = 0
	for {
		fmt.Scan(&warnabaju)
		if warnabaju == "." {
			break
		}
		A[i] = warnabaju
		i++
		size++
	}

	fmt.Scan(&searchKey)

	cekBaju(A, size, searchKey)
}

func cekBaju(A arrBaju, size int, searchKey string) {
	var found bool
	for i := 0; i < size; i++ {
		if A[i] == searchKey {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Baju Ada")
	} else {
		fmt.Println("Baju Tidak Ada")
	}
}
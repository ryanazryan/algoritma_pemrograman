package main

import "fmt"

const NMAX int = 5

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	if *n < NMAX {
		var x int
		fmt.Scan(&x)
		A[*n] = x
		*n++
	} else {
		fmt.Println("Array penuh")
	}
}	

func cetakData(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func main() {
	var data tabInt
	var nData int

	nData = 0
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	cetakData(data, nData)
	nData = 0
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	bacaData(&data, &nData)
	cetakData(data, nData)
}
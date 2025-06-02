package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main(){
	var NA tabInt
	var n int
	bacaNilai(&NA, &n)
	cetakNilai(NA, n)
}

func bacaNilai(NA *tabInt, n *int) {
	fmt.Print("Masukkan banyak elemen (n): ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		if i < NMAX {
			fmt.Printf("Masukkan elemen ke-%d: ", i+1)
			fmt.Scan(&(*NA)[i])
		} else {
			(*NA)[i] = NMAX
		}
	}
}

func cetakNilai(NA tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong.")
	} else {
		fmt.Print("Nilai yang terdapat pada array: ")
		for i := 0; i < n; i++ {
			fmt.Print(NA[i], " ")
		}
	}
}
package main

import "fmt"

const NMAX = 5

type tabInt struct {
	info [NMAX] int
	n int
}

func main() {
	var NA tabInt
	bacaNilai(&NA)
	cetakNilai(NA)
}

func bacaNilai(NA *tabInt) {
	fmt.Print("Masukkan banyak elemen (n): ")
	fmt.Scan(&NA.n)
	if NA.n > NMAX {
		NA.n = NMAX
	}

	fmt.Println("Masukkan nilai elemen: ")
	for i := 0; i < NA.n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&NA.info[i])
	}
}

func cetakNilai(NA tabInt) {
	if NA.n == 0 {
		fmt.Println("Array kosong")
		return
	}

	fmt.Println("Nilai yang terdapat pada array: ")
	for i := 0; i < NA.n; i++ {
		fmt.Print(NA.info[i], " ")
	}
	fmt.Println()
}
package main

import "fmt"

const NMAX = 20

type tabChar [NMAX] byte

func main() {
	var K tabChar
	var n int

	baca(&K, &n)
	cetak(K, n)
}

func baca(K *tabChar, n *int) {
	fmt.Print("Masukkan banyak elemen (n): ")
	fmt.Scan(n)

	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan karakter ke-%d: ", i+1)
		fmt.Scan(&(*K)[i])
	}
}

func cetak(k tabChar, n int) {
	if n == 0 {
		fmt.Println("Array kosong.")
		return
	}

	fmt.Print("Karakter yang terdapat pada array: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(string(k[i]), " ")
	}
	fmt.Println()
}
package main

import "fmt"

const NMAX int = 5

type matrix [NMAX][NMAX]int

func main() {
	var A, B, C matrix
	var n int

	baca(&A, &n)
	baca(&B, &n)

	fmt.Println("Matriks A:")
	cetak(A, n)
	fmt.Println("Matriks B:")
	cetak(B, n)

	jumlah(&A, &B, &C, n)

	fmt.Println("Matriks C (A + B):")
	cetak(C, n)
}

func baca(m *matrix, n *int) {
	fmt.Print("Masukkan ukuran matriks (n): ")
	fmt.Scan(n)

	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		for j := 0; j < *n; j++ {
			fmt.Printf("Masukkan elemen matriks [%d][%d]: ", i+1, j+1)
			fmt.Scan(&(*m)[i][j])
		}
	}
}

func cetak(m matrix, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
}

func jumlah(A, B, C *matrix, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			(*C)[i][j] = (*A)[i][j] + (*B)[i][j]
		}
	}
}
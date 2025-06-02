package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func main() {
	var A tabInt
	var n int
	baca(&A, &n)
	avg := rataRata(A, n)
	count := lebihDariRataRata(A, n, avg)
	fmt.Println("Rata-rata: ", avg)
	fmt.Println("Jumlah nilai lebih dari rata-rata: ", count)
}

func baca(A *tabInt, n *int) {
	fmt.Print("Masukkan jumlah elemen (n): ")
	fmt.Scan(n)

	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan nilai ke-%d: ", i+1)
		fmt.Scan(&(*A)[i])
	}
}

func rataRata(A tabInt, n int) float64 {
	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	return float64(sum) / float64(n)
}

func lebihDariRataRata(A tabInt, n int, avg float64) int {
	var count int
	for i := 0; i < n; i++ {
		if float64(A[i]) > avg {
			count++
		}
	}
	return count
}
package main

import "fmt"

const NMAX int = 20

type tFilm struct {
	title string
	year int
	review, rating float64
}

type tabFilm [NMAX] tFilm

func bacaDataFilm (A *tabFilm, n *int) {
	fmt.Scan(n)

	if *n > NMAX {
		*n = NMAX
	}

	var i int
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].title, &A[i].year, &A[i].review, &A[i].rating)
	}
}

func cetakDataFilm (A tabFilm, n int) {
	fmt.Printf("%41s %6s %6s %6s\n", "Title", "Year", "Review", "Rating")

	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%41s %d %6.2f %6.2f\n", A[i].title, A[i].year, A[i].review, A[i].rating)
	}
	fmt.Println()
}

func urutMenaik(A *tabFilm, n *int) {
	var temp tFilm
	var i, pass int 
	
	for pass = 1; pass < *n; pass++ {
		temp = A[pass]
		i = pass 
		for (i > 0 && A[i-1].year > temp.year) || (i > 0 && A[i].year == temp.year && A[i-1].review > temp.review){
			A[i] = A[i-1]
			i--
		}
		A[i] = temp

	}
}

func urutMenurun(A *tabFilm, n * int) {
	var i, idx, pass int
	var temp tFilm
	for pass = 1; pass < *n; pass++ {
		idx = pass-1
		for i = pass; i < *n; i++ {
			if (A[i].review > A[idx].review) || (A[i].review == A[idx].review && A[i].rating > A[idx].rating){
				idx = i
			}
		}

		 temp = A[idx]
		 A[idx] = A[pass-1]
		 A[pass-1] = temp
	}
}

func main() {
	var film tabFilm
	var nFilm int
	bacaDataFilm(&film, &nFilm)
	fmt.Println("Data sebelum diurutkan:")
	cetakDataFilm(film, nFilm)
	urutMenurun(&film, &nFilm)
	fmt.Println("Data setelah diurutkan menurun berdasarkan review dan rating:")
	cetakDataFilm(film, nFilm)
	urutMenaik(&film, &nFilm)
	fmt.Println("Data setelah diurutkan menaik berdasarkan year dan review")
	cetakDataFilm(film, nFilm)
}
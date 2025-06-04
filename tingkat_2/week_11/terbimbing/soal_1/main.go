package main

import "fmt"

const NMAX int = 1024

type books struct {
	bookID, title, author, genre string
	pubYear int
}

type tabBook [NMAX]books


func readData(A *tabBook, n *int){
	if *n > NMAX {
		*n = NMAX
	}

	var i int
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].bookID, &A[i].title, &A[i].author, &A[i].genre, &A[i].pubYear)
	}
}

func printData(A tabBook, n int) {
	fmt.Println("Daftar Buku: ")
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(A[i].bookID, A[i].title, A[i].author, A[i].genre, A[i].pubYear)
	}
}

func findData(A tabBook, n int, bID string) int {
	var idx, right, mid, left int
	idx = -1
	right = n- 1
	left = 0

	for left <= right &&  idx == -1 {
		mid = (left + right) / 2
		if A[mid].bookID == bID {
			idx = mid
		} else if A[mid].bookID < bID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return idx
}

func deleteData(A *tabBook, n *int, bID string) {
	idx := findData(*A, *n, bID)
	if idx != -1 {
		var i int
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n--
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func main() {
	var data tabBook
	var nData int
	var bookID string
	fmt.Scan(&nData)
	readData(&data, &nData)
	fmt.Scan(&bookID)
	deleteData(&data, &nData, bookID)
	printData(data, nData)
}
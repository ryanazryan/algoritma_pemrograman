package main

import "fmt"

const NMAX = 10

type employee struct {
	employeeID string
	name       string
	position   string
	department string
	joinYear   int
}


type tabEmployee [NMAX]employee

func readData(A *tabEmployee, n *int) {
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].employeeID, &A[i].name, &A[i].position, &A[i].department, &A[i].joinYear)
	}
}

func printData(A *tabEmployee, n int) {
	fmt.Println("Employee Data:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %s %d\n", A[i].employeeID, A[i].name, A[i].position, A[i].department, A[i].joinYear)
	}
}

func findData(A *tabEmployee, n int, ID string) int {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if A[mid].employeeID == ID {
			return mid
		} else if A[mid].employeeID < ID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func deleteData(A *tabEmployee, n *int, ID string) {
	idx := findData(A, *n, ID)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func main() {
	var data tabEmployee
	var nData int
	var bookID string

	fmt.Scan(&nData)
	readData(&data, &nData)
	fmt.Scan(&bookID)
	idx := findData(&data, nData, bookID)

	if idx != -1 {
		printData(&data, nData)
		deleteData(&data, &nData, bookID)
		printData(&data, nData)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
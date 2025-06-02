package main

import "fmt"

const maxSize = 100
type Arr [maxSize]int

func inputArray(arr *Arr, size *int, n int) {
	if *size+n > maxSize {
		fmt.Println("Input gagal, jumlah melebihi kapasitas maksimal")
		return
	} else {
		fmt.Println("Masukkan data:")

		var input int
		for i := 0; i < n; i++ {
			fmt.Printf("Data ke-%d: ", *size+1)
			fmt.Scan(&input)

			duplicate := false
			for j := 0; j < *size; j++ {
				if arr[j] == input {
					duplicate = true
					break
				}
			}

			if duplicate {
				fmt.Println("Data sudah ada, coba lagi.")
				i--
			} else {
				arr[*size] = input
				*size++
			}
		}
	}
}

func displayArray(arr Arr, size int) {
	fmt.Println("Isi Array:")
	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func main() {
	var arr Arr 
	var size int        
	var n int           

	fmt.Print("Masukkan jumlah data yang ingin dimasukkan: ")
	fmt.Scan(&n)
	inputArray(&arr, &size, n)
	displayArray(arr, size)
}
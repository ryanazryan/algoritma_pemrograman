package main

import "fmt"

const maxSize = 100
type Arr [maxSize]int

// poin a
func inputArray(arr *Arr, size *int, n int) {
	if *size + n > maxSize {
		fmt.Println("Input gagal, jumlah melebihi kapasitas maksimal")
		return
	} else {
		fmt.Println("Masukkan data")
		var i int
		for i = 0; i < n; i++ {
			fmt.Printf("Data ke-%d: ", i+1)
			fmt.Scan(&arr[*size])
			*size++
		}
	}
}

// poin b
func displayArray(arr Arr, size int) {
	fmt.Println("Isi Array:")
	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}


// poin c
func sumArray(arr Arr, size int) int {
	var sum int
	sum = 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	return sum
}

// poin d
func oddArray(arr Arr, size int) int {
	var count int = 0
	for i := 0; i < size; i++ {
		if arr[i] % 2 != 0 {
			count++
		}
	}
	return count
}

// poin e
func shiftArray(arr *Arr, size *int, i int, n int, j int) {
	if i < 0 || j < 0 || i + n > *size || j + n > *size {
		fmt.Println("Pergeseran tidak dapat dilakukan")
		return
	}
	var temp Arr
	var k int
	for k = 0; k < n; k++ {
		temp[k] = arr[i+k]
	}
	for k = i + n; k < *size; k++ {
		arr[k-n] = arr[k]
	}
	*size -= n
	for k := *size - 1; k >= j; k-- {
		arr[k+n] = arr[k]
	}
	for k = 0; k < n; k++ {
		arr[j+k] = temp[k]
	}
	*size += n
}

// poin f
func reverseArray(arr *Arr, size int) Arr {
	var reversed Arr
	for i := 0; i < size; i++ {
		reversed[i] = (*arr)[size-1-i]
	}
	return reversed
}

// poin g
func mergeSortedArrays(arr1, arr2 Arr, size1, size2 int) Arr {
	var result Arr
	i, j, k := 0, 0, 0
	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	for i < size1 {
		result[k] = arr1[i]
		i++
		k++
	}
	for j < size2 {
		result[k] = arr2[j]
		j++
		k++
	}
	return result
}

func main() {
	var arr1, arr2 Arr
	var size1, size2 int

	fmt.Println("Masukkan data untuk Array 1:")
	inputArray(&arr1, &size1, 5)
	fmt.Println("Masukkan data untuk Array 2:")
	inputArray(&arr2, &size2, 5)
	fmt.Println("Array 1:")
	displayArray(arr1, size1)
	fmt.Println("Array 2:")
	displayArray(arr2, size2)
	total := sumArray(arr1, size1)
	fmt.Println("Total jumlah Array 1:", total)
	oddCount := oddArray(arr1, size1)
	fmt.Println("Jumlah bilangan ganjil pada Array 1:", oddCount)
	reversed := reverseArray(&arr1, size1)
	fmt.Println("Array 1 setelah dibalik:")
	displayArray(reversed, size1)
	mergedArray := mergeSortedArrays(arr1, arr2, size1, size2)
	fmt.Println("Array yang sudah digabungkan dan terurut:")
	displayArray(mergedArray, size1+size2)
	fmt.Println("Array sebelum pergeseran:")
	displayArray(arr1, size1)
	shiftArray(&arr1, &size1, 2, 2, 4)
	fmt.Println("Array setelah pergeseran:")
	displayArray(arr1, size1)
}
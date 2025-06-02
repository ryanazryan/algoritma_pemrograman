package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Print("Masukkan nilai m: ")
	fmt.Scan(&m)
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	deret := make([]int, n)
	fmt.Println("Masukkan deret bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&deret[i])
	}

	valid := cekPerkalian(m, n, deret)
	fmt.Println("Apakah deret bilangan merupakan hasil perkalian?", valid)
}

func cekPerkalian(m, n int, deret []int) bool {
	for i := 1; i <= n; i++ {
		expected := m * i
		if deret[i-1] != expected {
			return false
		}
	}
	return true
}

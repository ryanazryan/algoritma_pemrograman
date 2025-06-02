package main

import (
	"fmt"
	"math"
	"sort"
)

func quadraticRoots(a, b, c float64) (float64, float64) {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		fmt.Println("Persamaan tidak memiliki akar real.")
		return math.NaN(), math.NaN()
	}

	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	roots := []float64{root1, root2}
	sort.Float64s(roots)

	return roots[0], roots[1]
}

func main2() {
	var a, b, c float64

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	root1, root2 := quadraticRoots(a, b, c)

	if !math.IsNaN(root1) && !math.IsNaN(root2) {
		fmt.Printf("Akar-akarnya adalah: %.3f dan %.3f\n", root1, root2)
	}
}
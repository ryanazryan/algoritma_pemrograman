package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Println("Masukkan ukuran papan (n) dan posisi awal Gajah (x, y):")
	fmt.Scan(&n, &x, &y)

	fmt.Println("Langkah Gajah yang dapat dijangkau:")
	for i := 1; i < n; i++ {
		if x+i <= n && y+i <= n {
			fmt.Printf("(%d, %d)\n", x+i, y+i)
		}
		if x-i > 0 && y+i <= n {
			fmt.Printf("(%d, %d)\n", x-i, y+i)
		}
		if x+i <= n && y-i > 0 {
			fmt.Printf("(%d, %d)\n", x+i, y-i)
		}
		if x-i > 0 && y-i > 0 {
			fmt.Printf("(%d, %d)\n", x-i, y-i)
		}
	}
}

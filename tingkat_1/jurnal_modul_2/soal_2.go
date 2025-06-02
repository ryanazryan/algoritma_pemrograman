package main

import "fmt"

func main2(){
	var tahunLahir int
	const tahunSekarang = 2024

	fmt.Print("Masukkan tahun lahir: ")
	fmt.Scan(&tahunLahir)

	usia := tahunSekarang - tahunLahir

	fmt.Printf("Teman Kuromi yang lahir pada tahun %d\n", tahunLahir)
	fmt.Printf("Saat ini telah berusia %d tahun\n", usia)

	// Penjelasan singkat
	// Rumus Usia: usia = tahun sekarang - tahun lahir
}
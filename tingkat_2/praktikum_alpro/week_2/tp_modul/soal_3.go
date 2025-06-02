package main

import "fmt"

func main()  {
	type product struct {
		nama string
		harga float64
		stok int
	}

	var p product

	for {
		fmt.Scan(&p.nama)
		if p.nama == "#" {
			break
		}

		fmt.Scan(&p.harga, &p.stok)

		var diskon float64
		if p.harga >= 1000000 {
			diskon = 0.10
		} else if p.harga >= 500000 {
			diskon = 0.05
		} else {
			diskon = 0.0
		}

		hargaSetelahDiskon := p.harga * (1 - diskon)

		status := "Produk habis"
		if p.stok > 0 {
			status ="Produk tersedia"
		}

		fmt.Println("Detail Produk: ")
		fmt.Println("Nama: ", p.nama)
		fmt.Printf("Harga Asli: Rp.%.0f\n", p.harga)
		fmt.Printf("Diskon: %.0f%%\n", diskon*100)
		fmt.Printf("Harga Setelah Diskon: Rp.%.0f\n", hargaSetelahDiskon)
		fmt.Println("Stok: ", p.stok)
		fmt.Println("Status: ", status)
		fmt.Println()
	}
}
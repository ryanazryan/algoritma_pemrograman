package main

import "fmt"

const NMAX int = 1024

type infoBuku struct {
	idBuku, judulBuku, penulisBuku string
}

type arrBuku [NMAX]infoBuku

func main() {
	var dataBuku arrBuku
	var nBuku int
	var searchKey string
	fmt.Scan(&nBuku)
	for i := 0; i < nBuku; i++ {
		fmt.Scan(&dataBuku[i].idBuku, &dataBuku[i].judulBuku, &dataBuku[i].penulisBuku)
	}
	fmt.Scan(&searchKey)
	pencarianBuku(dataBuku, nBuku, searchKey)
}

func pencarianBuku(dataBuku arrBuku, nBuku int, searchKey string) {
	found := false
	for i := 0; i < nBuku; i++ {
		if dataBuku[i].judulBuku == searchKey {
			fmt.Printf("ID Buku: %s\n", dataBuku[i].idBuku)
			fmt.Printf("Judul Buku: %s\n", dataBuku[i].judulBuku)
			fmt.Printf("Penulis Buku: %s\n", dataBuku[i].penulisBuku)
			found = true
			break
		}
	}
	
	if !found {
		fmt.Println("Buku Tidak Tersedia")
	}
}
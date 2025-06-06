package main

import "fmt"

const NMAX int = 5

type buku struct {
	id, judul string
	tahun, halaman int
}

type tabBuku [NMAX]buku

func tampilkanMenu (pilih *int) {
	fmt.Println("----------------------------------------")
	fmt.Println("                MENU")
	fmt.Println("----------------------------------------")
	fmt.Println("1. Input Buku")
	fmt.Println("2. Output Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Cari Buku Tahun Terbaru")
	fmt.Println("5. Cari Buku Berdasarkan ID")
	fmt.Println("6. Urut Buku Menaik Berdasarkan Jumlah Halaman")
	fmt.Println("7. Cari Buku Berdasarkan Tahun")
	fmt.Println("8. Urut Buku Menurun Berdasarkan Tahun")
	fmt.Println("9. Exit")
	fmt.Println("----------------------------------------")
	fmt.Print("Pilih (1-9): ")
	fmt.Scan(pilih)
}

func inputBuku(A *tabBuku, n *int) {
	var id, judul string
	var tahun, halaman int
	fmt.Scan(&id, &judul, &tahun, &halaman)
	if *n < NMAX {
		A[*n].id, A[*n].judul, A[*n].tahun, A[*n].halaman = id, judul, tahun, halaman
		*n++
		fmt.Println("Penambahan berhasil, data bertambah 1")
	} else {
		fmt.Println("Penambahan gagal, data penuh")
	}

}

func outputBuku(A tabBuku, n int) {
	if n!= 0 {
		var i int
		for i = 0; i < n; i++ {
			fmt.Printf("%s %s %d %d\n", A[i].id, A[i].judul, A[i].tahun, A[i].halaman)
		} 
	} else {
		fmt.Println("Data Kosong")
	}
}

func hapusBuku(A *tabBuku, n *int) {
	if *n > 0 {
		*n--
		fmt.Println("Penghapusan berhasil, data berkurang 1")
	} else {
		fmt.Println("Penghapusan gagal")
	}
}

func cariBukuTerbaru(A tabBuku, n int) string{
	var i, idx int
	idx = 0
	i = 1
	for i < n{
		if A[i].tahun > A[idx].tahun {
			idx = i
		}
		i++
	}
	return A[idx].judul
}

func cariBukuBerdasarkanID(A tabBuku, n int, id string) string {
	var i, idx int
	idx = -1
	i = 0

	for i < n && idx == -1 {
		if A[i].id == id{
			idx = i
		}
		i++
	}

	if idx != -1 {
		return A[idx].judul
	} else {
		return "Buku tidak ditemukan"
	}
}

func urutBukuMenaik(A *tabBuku, n int){
	var i, idx, pass int
	var temp buku

	for pass = 1; pass < n; pass++{
		idx = pass -1
		for i = pass; i < n; i++ {
			if A[i].halaman < A[idx].halaman{
				idx = i
			}
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
}

func cariBukuBerdasarkanTahun(A tabBuku, n, tahun int) string {
	var left, right, mid int
	var idx int

	idx = -1
	left = 0
	right = n-1

	for left <= right && idx == -1 {
		mid = (left + right / 2)
		if A[mid].tahun == tahun {
			idx = mid
		} else if A[mid].tahun > tahun{
			left = mid + 1
		} else {
			right = mid -1
		}
	}

	if idx!= -1 {
		return A[idx].judul
	} else {
		return "Buku tidak ditemukan"
	}
}

func urutBukuMenurun(A *tabBuku, n int) {
	var i, pass int
	var temp buku

	for pass = 1; pass < n; pass++ {
		i = pass 
		temp = A[pass]
		for i > 0 && A[i-1].tahun < temp.tahun {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
	}
}

func main() {
	var data tabBuku
	var nData int
	var pilihan, tahun int
	var id string

	for {
		tampilkanMenu(&pilihan)
		switch pilihan {
		case 1:
			inputBuku(&data, &nData)
		case 2:
			outputBuku(data, nData)
		case 3:
			hapusBuku(&data, &nData)
		case 4:
			fmt.Println(cariBukuTerbaru(data, nData))
		case 5:
			fmt.Print("Masukkan ID buku yang dicari:")
			fmt.Scan(&id)
			fmt.Println(cariBukuBerdasarkanID(data, nData, id))
		case 6:
			urutBukuMenaik(&data, nData)
		case 7:
			fmt.Print("Masukkan tahun buku yang dicari:")
			fmt.Scan(&tahun)
			fmt.Println(cariBukuBerdasarkanTahun(data, nData, tahun))
		case 8:
			urutBukuMenurun(&data, nData)
		}
		if pilihan == 9 {
			break
		}
	}
}
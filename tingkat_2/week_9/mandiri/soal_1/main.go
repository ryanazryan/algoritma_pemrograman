package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol [NMAX]int
	nama string
	totalPertandingan, totGol, totKebobolan int
	totMenang, totKalah, totDraw, totPoint int
}

func main() {
	var timA, timB tim
	bacaHasil(&timA, &timB)
	fmt.Println("Statistik Tim A:", timA.nama)
	cetakHasil(timA)
	fmt.Println("Statistik Tim B:", timB.nama)
	cetakHasil(timB)
}

func bacaHasil(tA, tB *tim) {
	var x, y int
	fmt.Print("Masukkan skor pertandingan tim A dan B (masukkan 0 0 untuk berhenti):\n")
	for {
		fmt.Scan(&x, &y)
		if x == 0 && y == 0 {
			break
		}
		tA.gol[tA.totalPertandingan] = x
		tB.gol[tB.totalPertandingan] = y
		tA.totGol += x
		tB.totGol += y
		tA.totKebobolan += y
		tB.totKebobolan += x

		if x > y {
			tA.totPoint += 3
			tA.totMenang++
			tB.totKalah++
		} else if x == y {
			tA.totPoint++
			tB.totPoint++
			tA.totDraw++
			tB.totDraw++
		} else {
			tB.totPoint += 3
			tB.totMenang++
			tA.totKalah++
		}

		tA.totalPertandingan++
		tB.totalPertandingan++
	}
}

func cetakHasil(t tim) {
	var n int = t.totalPertandingan
	fmt.Println("Total Pertandingan:", n)
	fmt.Print("Gol tiap pertandingan: ")
	for i := 0; i < n; i++ {
		fmt.Print(t.gol[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	fmt.Println("Total Menang:", t.totMenang)
	fmt.Println("Total Draw:", t.totDraw)
	fmt.Println("Total Kalah:", t.totKalah)
	fmt.Println("Total Gol:", t.totGol)
	fmt.Println("Total Kebobolan:", t.totKebobolan)
	fmt.Println("Total Point:", t.totPoint)
}
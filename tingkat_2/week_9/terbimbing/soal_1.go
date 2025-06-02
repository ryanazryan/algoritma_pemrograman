package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol [NMAX] int
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

func bacaHasil(tA, tB * tim) {
	var x, y int
	fmt.Scan(&x, &y)
	var i int = 0
	for x > 0 && y > 0 && i < NMAX {
		tA.gol[i] = x
		tB.gol[i] = y
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
		i++
		fmt.Scan(&x, &y)
	}

	tA.totalPertandingan = i
	tB.totalPertandingan = i
}

func cetakHasil(t tim) {
	var n int = t.totalPertandingan
	fmt.Println("Total Pertandingan:", n)
	fmt.Print("Gol tiap pertandingan: ")
	var i int
	for i = 0; i < n; i++ {
		fmt.Print(t.gol[i], "")
		if i < n-1 {
			fmt.Print(t.gol[i], " ")
		}
		fmt.Println()
		fmt.Println("Total Menang:", t.totMenang)
		fmt.Println("Total Draw:", t.totDraw)
		fmt.Println("Total Kalah:", t.totKalah)
		fmt.Println("Total Gol:", t.totGol)
		fmt.Println("Total Kebobolan:", t.totKebobolan)
		fmt.Println("Total Point:", t.totPoint)
	}
}
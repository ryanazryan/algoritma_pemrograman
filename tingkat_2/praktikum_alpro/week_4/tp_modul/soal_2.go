package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg += g - k
}

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = (jm * 3) + (jd * 1)
}

func main2() {
	var N int
	fmt.Scan(&N)

	jm, jd, jk := 0, 0, 0
	jg, jkg, jsg := 0, 0, 0
	jp := 0

	for i := 0; i < N; i++ {
		var g, k int
		fmt.Scan(&g, &k)

		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)
		hitungJumGolKegolanSelisih(g, k, &jg, &jkg, &jsg)
	}
	hitungJumPoint(jm, jd, &jp)

	fmt.Printf("%d %d %d %d %d %d %d %d\n", N, jm, jd, jk, jg, jkg, jsg, jp)
}
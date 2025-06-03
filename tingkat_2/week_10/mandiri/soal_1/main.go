package main

import "fmt"

const NMAX int = 1024

type waktu struct {
	menit, detik int
}

type arrWaktu [NMAX]waktu

func main() {
	var waktuPembalap arrWaktu
	var nPembalap, waktuMinimal int
	var menitMulai, detikMulai int
	var i int
	fmt.Scan(&nPembalap)
	fmt.Scan(&menitMulai, &detikMulai)
	for i = 0; i < nPembalap; i++ {
		fmt.Scan(&waktuPembalap[i].menit, &waktuPembalap[i].detik)
	}
	waktuMinimal = waktuTercepat(waktuPembalap, nPembalap, detikMulai, menitMulai)
	fmt.Println(waktuMinimal/60, waktuMinimal%60)
}

func convertWaktu(waktuMenit int, waktuDetik int) int {
	detik := waktuMenit*60 + waktuDetik
	return detik
}

func waktuTercepat(waktuPembalap arrWaktu, nPembalap, detikMulai int, menitMulai int) int {
	var i int
	var min int = 999999999
	var cur int
	var start int = convertWaktu(menitMulai, detikMulai)
	for i = 0; i < nPembalap; i++ {
		cur = convertWaktu(waktuPembalap[i].menit, waktuPembalap[i].detik) - start
		if min > cur {
			min = cur
		}
	}
	return min
}
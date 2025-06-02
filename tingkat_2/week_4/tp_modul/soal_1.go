package main

import "fmt"

const pi = 3.14

func hitungLuasKelilingLingkaran(r float64, l *float64, k *float64) {
	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l *float64, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, totLuas *float64, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}

func main1() {
	var r, s float64
	var Rs, Ss, LLs, LPs, KLs, KPs, TLs, TPs []float64

	for {
		fmt.Scan(&r, &s)
		if r == 0 && s == 0 {
			break
		}

		var lL, lP, kL, kP, totLuas, totKel float64

		hitungLuasKelilingLingkaran(r, &lL, &kL)
		hitungLuasKelilingPersegi(s, &lP, &kP)
		hitungTotal(lL, lP, kL, kP, &totLuas, &totKel)

		Rs = append(Rs, r)
		Ss = append(Ss, s)
		LLs = append(LLs, lL)
		LPs = append(LPs, lP)
		KLs = append(KLs, kL)
		KPs = append(KPs, kP)
		TLs = append(TLs, totLuas)
		TPs = append(TPs, totKel)
	}

	fmt.Printf("%-10s%-10s%-10s%-10s%-10s%-10s%-10s%-10s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")

	for i := 0; i < len(Rs); i++ {
		fmt.Printf("%-10.2f%-10.2f%-10.2f%-10.2f%-10.2f%-10.2f%-10.2f%-10.2f\n",
			Rs[i], Ss[i], LLs[i], LPs[i], KLs[i], KPs[i], TLs[i], TPs[i])
	}
}

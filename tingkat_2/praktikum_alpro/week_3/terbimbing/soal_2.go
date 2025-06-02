package main

import "fmt"

func ganjil(a, b, c, d, e int) bool {
	var cek bool
	cek = a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0
	return cek
}

func genap(a,b,c,d, e int ) bool {
	var cek bool 
	cek = a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0
	return cek
}

func diskon(member bool, c, d, e int ) float64 {
	var disc float64
	disc = 1.5*float64(c+d+e)
	if member {
		disc = disc + 0.25*disc
	}

	if disc > 50{
		disc = 50
	}

	return disc
}

func cashback(member bool, a, b, c int) float64 {
	var cash float64
	cash = 2.5*float64(a+b+c)
	if member {
		cash = cash + 0.15*cash
	}

	if cash > 40 {
		cash = 40
	}

	return cash
}

func main2(){
	var m bool
	var a, b, c, d, e int
	var disc float64
	var cash float64
	fmt.Scan(&m, &a, &b, &c, &d, &e)
	if ganjil(a, b, c, d, e){
		disc = diskon(m, c, d, e)
	} else if genap(a, b, c, d, e){
		cash = cashback(m, a, b, c)
	} else {
		disc = diskon(m, c, d, e)
		cash = cashback(m, a, b, c)
	}

	fmt.Printf("%.3f %.3f", disc, cash)
}
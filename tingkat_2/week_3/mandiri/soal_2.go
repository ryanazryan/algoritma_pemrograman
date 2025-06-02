package main

import (
	"fmt"
	"math"
)

func konversiDerajatkeRadian(T float64) float64 {
	return T * math.Pi / 180
}

func waktu(V, T float64) float64 {
	g := 9.8
	T_radian := konversiDerajatkeRadian(T)
	return V * math.Sin(T_radian) / g
}

func jarak(V, T float64) float64 {
	g := 9.8
	T_radian := konversiDerajatkeRadian(2 * T)
	return (V * V * math.Sin(T_radian)) / g
}

func ketinggian(V, T float64) float64 {
	g := 9.8
	T_radian := konversiDerajatkeRadian(T)
	return (V * V * math.Sin(T_radian) * math.Sin(T_radian)) / (2 * g)
}

func main() {
	var V, T float64
	fmt.Scanln(&V, &T)

	waktuTempuh := waktu(V, T)
	jarakHorizontal := jarak(V, T)
	ketinggianMaksimum := ketinggian(V, T)

	fmt.Printf("%.2f %.2f %.2f\n", waktuTempuh, jarakHorizontal, ketinggianMaksimum)
}
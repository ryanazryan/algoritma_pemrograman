package main

import "fmt"

func fungsi_f(x, y, z float64) float64 {
	return (2 * x) / (x + y + z)
}

func fungsi_g(x, y float64) float64 {
	return (1 / (5 * x)) + y
}

func main1(){
	var b1, b2, b3 float64
	var result_f, result_g float64

	fmt.Scanf("%f %f %f", &b1, &b2, &b3)

	result_f = fungsi_f(b1, b2, b3)
	result_g = fungsi_g(b1, b2)

	fmt.Printf("%.15f %.1f\n", result_f, result_g)
}
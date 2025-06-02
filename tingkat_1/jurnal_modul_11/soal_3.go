package main

import "fmt"

func main() {
	var gravitasi float64
	var suhu int
	var atmosfer string

	fmt.Println("Masukkan nilai gravitasi, suhu, dan keberadaan atmosfer (contoh: 1.0 20 ada):")
	fmt.Scan(&gravitasi, &suhu, &atmosfer)

	switch {
	case gravitasi < 0.5 || gravitasi > 1.5:
		if suhu < -50 || suhu > 50 {
			if atmosfer != "ada" {
				fmt.Println("Pendaratan tidak memungkinkan: gravitasi, suhu, dan atmosfer planet tidak aman.")
			} else {
				fmt.Println("Pendaratan tidak memungkinkan: gravitasi dan suhu planet tidak aman.")
			}
		} else if atmosfer != "ada" {
			fmt.Println("Pendaratan tidak memungkinkan: gravitasi dan atmosfer planet tidak aman.")
		} else {
			fmt.Println("Pendaratan tidak memungkinkan: gravitasi planet tidak aman.")
		}
	case suhu < -50 || suhu > 50:
		if atmosfer != "ada" {
			fmt.Println("Pendaratan tidak memungkinkan: suhu dan atmosfer planet tidak aman.")
		} else {
			fmt.Println("Pendaratan tidak memungkinkan: suhu planet tidak aman.")
		}
	case atmosfer != "ada":
		fmt.Println("Pendaratan tidak memungkinkan: atmosfer planet tidak aman.")
	default:
		fmt.Println("Pendaratan memungkinkan: kondisi planet aman untuk pendaratan.")
	}
}

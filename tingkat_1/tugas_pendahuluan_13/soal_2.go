package main

import (
	"fmt"
)

func main() {
	var kekuatan, totalKarakter, totalKekuatan int
	kekuatanTertinggi := -1
	kekuatanTerendah := 101

	for {
		fmt.Print("Masukkan kekuatan karakter (1-99, 100 untuk berhenti): ")
		fmt.Scan(&kekuatan)

		if kekuatan == 100 {
			break
		}

		if kekuatan < 1 || kekuatan > 99 {
			continue
		}

		totalKarakter++
		totalKekuatan += kekuatan

		if kekuatan > kekuatanTertinggi {
			kekuatanTertinggi = kekuatan
		}

		if kekuatan < kekuatanTerendah {
			kekuatanTerendah = kekuatan
		}
	}

	if totalKarakter > 0 {
		rataRata := float64(totalKekuatan) / float64(totalKarakter)
		fmt.Printf("Total karakter: %d\n", totalKarakter)
		fmt.Printf("Rata-rata kekuatan: %.2f\n", rataRata)
		fmt.Printf("Kekuatan terendah: %d\n", kekuatanTerendah)
		fmt.Printf("Kekuatan tertinggi: %d\n", kekuatanTertinggi)
	} else {
		fmt.Println("Tidak ada data karakter yang valid.")
	}
}

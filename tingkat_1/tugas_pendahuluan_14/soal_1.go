package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah mobil: ")
	jumlahMobilStr, _ := reader.ReadString('\n')
	jumlahMobilStr = strings.TrimSpace(jumlahMobilStr)

	jumlahMobil := 0
	fmt.Sscanf(jumlahMobilStr, "%d", &jumlahMobil)

	fmt.Print("Masukkan nama tim 1: ")
	tim1, _ := reader.ReadString('\n')
	tim1 = strings.TrimSpace(tim1)

	fmt.Print("Masukkan nama tim 2: ")
	tim2, _ := reader.ReadString('\n')
	tim2 = strings.TrimSpace(tim2)

	pointMap := map[int]int{
		1: 7,
		2: 5,
		3: 3,
		4: 1,
		5: 0,
		6: 0,
	}

	fmt.Println("Masukkan urutan finish masing-masing mobil (tim 1 atau tim 2):")
	tim1Points := 0
	tim2Points := 0

	for i := 1; i <= jumlahMobil; i++ {
		fmt.Printf("Mobil ke-%d (tim 1 atau tim 2): ", i)
		finish, _ := reader.ReadString('\n')
		finish = strings.TrimSpace(finish)

		if finish == tim1 {
			if i <= 6 {
				tim1Points += pointMap[i]
			}
		} else if finish == tim2 {
			if i <= 6 {
				tim2Points += pointMap[i]
			}
		} else {
			fmt.Println("Input tim tidak valid.")
			return
		}
	}

	fmt.Printf("%s memperoleh poin %d\n", tim1, tim1Points)
	fmt.Printf("%s memperoleh poin %d\n", tim2, tim2Points)

	if tim1Points > tim2Points {
		fmt.Printf("Pemenang adalah tim %s\n", tim1)
	} else if tim2Points > tim1Points {
		fmt.Printf("Pemenang adalah tim %s\n", tim2)
	} else {
		fmt.Println("Hasil seri!")
	}
}

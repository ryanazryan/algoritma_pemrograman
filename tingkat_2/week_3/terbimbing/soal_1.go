package main
import "fmt"

func numeric(karakter byte) bool{
	var cek bool
	cek = karakter >= '0' && karakter <= '9'
	return cek
}

func main () {
	var kar byte
	var count int
	count = 0

	fmt.Scanf("%c", &kar)
	for kar != '.' {
		if numeric(kar) {
			count++
		}
		fmt.Scanf("%c", &kar)
	}

	fmt.Print("Jumlah karakter numerik adalah ", count)
}
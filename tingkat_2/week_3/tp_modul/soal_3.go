package main
import "fmt"

func adaBilanganM(n, m int) string {
	for n > 0 {
		digit := n % 10
		if digit == m {
			return "YA"
		}

		n = n / 10
	}

	return "TIDAK"
}

func main(){
	var n, m int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Println(adaBilanganM(n, m))
}
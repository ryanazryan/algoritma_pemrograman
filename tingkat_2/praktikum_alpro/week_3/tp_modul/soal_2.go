package main
import "fmt"

func lowToUpper(karakter byte) byte{
	if karakter >= 'a' && karakter <= 'z' {
		return karakter - 'a' + 'A'
	}
	return karakter
}

func main2(){
	var kar byte
	fmt.Scanf("%c\n", &kar)
	fmt.Printf("%c", lowToUpper(kar))
}
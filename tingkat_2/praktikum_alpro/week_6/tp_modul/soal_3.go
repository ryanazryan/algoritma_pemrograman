package main

import "fmt"

const ukuranUbin int = 30 * 30

func conversionToCm(length, width int) (int, int) {
	length = length * 100
	width = width * 100
	return length, width
}

func calRoom(length, width int) int {
	return length * width
}

func countTiles(roomArea int) int {
	tilesNeeded := roomArea / ukuranUbin
	if roomArea % ukuranUbin != 0 {
		tilesNeeded ++
	}

	return tilesNeeded
}

func main(){
	var L, W int
	fmt.Scan(&L, &W)

	L, W = conversionToCm(L, W)
	roomArea := calRoom(L, W)
	tiles := countTiles(roomArea)

	fmt.Printf("You need: %d tiles\n", tiles)
}
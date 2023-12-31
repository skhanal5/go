package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := range slice {
		slice[i] = make([]uint8, dx)
	}

	for i := range slice {
		val := slice[i]
		for j := range val {
			slice[i][j] = 1
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}

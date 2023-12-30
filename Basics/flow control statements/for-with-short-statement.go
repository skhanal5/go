package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	/*
		Can provide an initializer
		that has block scope
	*/
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {

	// Expected 9 and 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i:=0; i<10; i++ {
		curr :=  (z*z - x) / (2*z)
		if z == curr {
			break
		} else {
			z -= curr
		}
	}

	return z
}

func main() {  
	fmt.Println(Sqrt(2))
}

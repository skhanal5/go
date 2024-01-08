package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f",float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0

	if x < 0.0 {
		a := ErrNegativeSqrt(x)
		return x, a
	}

	for i:=0; i<10; i++ {
		curr :=  (z*z - x) / (2*z)
		if z == curr {
			break
		} else {
			z -= curr
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
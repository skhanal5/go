package main

import (
	"fmt"
	"math"
)

/*
	You can declare methods on non-struct types too
	Here we do it on a numeric type                               

	Note: You can only declare a method w/ a receiver
	whose type is defined in the SAME PACKAGE

	For example: we cannot use int as our receiver type
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

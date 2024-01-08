package main

import (
	"fmt"
	"math/cmplx"
)


/*
	1. Variable declarations can be factored like import statements
	2. Types:
		bool, string, int[8-64], uint[8-64, ptr], byte, rune, 
		float(32/64), complex(32/64)
		[] indicates optional values after keyword
		() indicates that the value is needed
*/

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main() {
	fmt.Printf("Type: %T Value; %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z,z)
}
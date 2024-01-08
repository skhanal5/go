package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	/*
		hypot is a function that computes the hypotenuse
		think: a^2 + b^2
	*/
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// We then invoke hypot
	fmt.Println(hypot(5, 12))


	// Then we pass in hypot into compute as a parameter
	// compute invokes hypot(3,4)
	fmt.Println(compute(hypot))

	// We then pass in math.pow into compute
	// compute invokes math.pow(3,4)
	fmt.Println(compute(math.Pow))
}

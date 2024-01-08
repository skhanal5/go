package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
	To summarize:
	We use method pointers for two reasons
	
	1. We can modify the value the receiver points to
	2. We can avoid copying the value on each method call -> inefficient

	Note: A good practice is to avoid having methods of a given type
	from mixing and matching value or pointer receivers
*/


/*
	This is an interesting example.
	We see a METHOD with a POINTER receiver
	and a VALUE argument.
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

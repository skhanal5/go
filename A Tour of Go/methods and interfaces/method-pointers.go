package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	A couple of things to make note here:
	
	- We can have pointer receivers in the form *T where
	T is a type defined in our package
	- methods with pointer receivers can modify the underyling
	value
	- this is a COMMON pattern where we use pointer receivers
	in place of value receivers

	Why?:
	- Methods with a value receiver use a copy of the original value

	See what happens when you change Scale from a pointer receiver
	to a value receiver.
*/

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

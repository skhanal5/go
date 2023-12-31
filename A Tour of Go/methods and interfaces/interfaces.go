package main

import (
	"fmt"
	"math"
)

/*
	An interface is a set of method signatures (like Java)
	A value of interface type can hold any value that
	implements those methods

	Interface types are strict
	Say you have a new type and it doesn't implement the
	method signature of the interface to a tee, then
	we cannot create a interface type and assign it to the
	new type. It violates the interface contract.
*/

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

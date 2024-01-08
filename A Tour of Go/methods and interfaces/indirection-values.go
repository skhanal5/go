package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}


/*
	Unlike the Scale() method, our Abs()
	method takes in a Vertex VALUE.

	What happens if we pass in a Vertex pointer?
	- It still works. This is part of the functionality
	shown in indirections.go
	- METHODS will values receivers will take POINTERS
	or VALUES as the receiver

*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	Here AbsFunc expects a Vertex value.
	If we pass in a pointer, it obviously fails
	like we expect it to.

	FUNCTIONS are still strict with argument types.
*/
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

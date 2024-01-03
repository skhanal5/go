package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

/*
	Method syntax differs from function syntaxc

	method syntax: func (receiver) name type
	function syntax: func name (args) type

	A receiver is a special type of argument

	Note: you can define methods on types (here it is Vertex)
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y) 
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
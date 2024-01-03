package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

/*
	Interface implementation for
	type T

	Note: t is a pointer
*/
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

/*
	Interface implementation for
	type F

	Note: f is a value
*/
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I		//i is of type interface I

	i = &T{"Hello"}
	describe(i)
	i.M()		// this works because T implements M with a pointer arg

	i = F(math.Pi)
	describe(i)
	i.M()		// this works becasue F implements M with a value arg
}

/*
	Interface values can be thought of like this:
	(value, type)

*/
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
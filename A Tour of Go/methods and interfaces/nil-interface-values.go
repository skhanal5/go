package main

import "fmt"


/*
	A nil interface doesn't hold a value or a type.
	
	If you call a method using a Nil receiver, you
	get a run-time error since we don't have a type
	that can be used to say "run this function"
*/
type I interface {
	M()
}

func main() {
	var i I		// I is not defined
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

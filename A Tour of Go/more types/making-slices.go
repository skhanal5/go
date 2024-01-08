package main

import "fmt"


/*
	This is a little confusing....

	make allocates an array filled with 0's and returns a slice
	- the first argument is the length of the slice
	- the optional second argument is the size of the capacity
*/
func main() {

	// a should be [0,0,0,0,0]
	a := make([]int, 5)
	printSlice("a", a)
	
	// b should be [] but w/ cap = 5
	b := make([]int, 0, 5)
	printSlice("b",b)

	// c is 0,0
	c := b[:2]
	printSlice("c",c)

	// d is [0,0,0]. This is the weird one, our cap changes..
	// i think this implies that slices are passed by value
	d := c[2:5]
	printSlice("d",d)

	// printSlice("c", c) 		c is unchanged here
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

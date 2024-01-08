package main

import "fmt"


/*
	Here we try out the append function.

	Append takes in two (or more) arguments:
	- arg 1: a slice of type T
	- arg 2 - n: arguments of type T that will be appended onto the slice

	If the slice is too small to contain the additional arguments,
	it adjusts the size of the slice behind the scenes to make it work
	by mkaing a new array and giving us a reference to it
*/

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s,0)
	printSlice(s)

	// we can add more than one element at a time
	s = append(s,2,3,4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

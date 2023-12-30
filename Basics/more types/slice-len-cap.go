package main

import "fmt"

/*
	Each slice has two fields:
	- a length (number of elements a slice has)
	- a capacity (comes from the underlying array)
		- this is counted from the first element in the slice

*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4=]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]		// We cut off the first two values, so our cap is 4
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

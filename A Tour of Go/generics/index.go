package main

import "fmt"

/*
	IMPORTANT NOTE:
	Go functions can work on multiple types using type parameters

					    -type params-
	syntax: func Index[T comparable](s []T, x T) int

	This is saying s is a slice of type T and x is of type T that is 
	also comparable
	
	comparable: makes it easy to use == and != operators on 
	values of the type
*/


// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
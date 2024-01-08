package main

import "fmt"

func main() {

	/*
		This creates a slice literal

		A slice literal is an array literal
		but we don't specify the length.

		Syntax: []type

		What is happening under the hood?

		1. An Array is being created w/ a specified size
		and any initialied contents
		2. The slice is referencing that Array
		3. End-user doesn't know about the Array (technically)
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// Note: we create a slice literaly of type struct where the struct has a int and bool field
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

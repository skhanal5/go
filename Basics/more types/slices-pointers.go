package main

import "fmt"


/*
	Slices do not actually store any data.
	Slices are used to DESCRIBE an existing Array


	You change anything in a slice, you end up changing
	the Array itself

	Other slices that derive from the same Array will be
	updated to reflect the changes

*/
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// a -> [John, Paul]
	// b -> [Paul, George]
	// we are replacing Paul with XXX in BOTH instances AND the array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

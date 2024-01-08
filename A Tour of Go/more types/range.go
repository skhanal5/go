package main

import "fmt"

var pow = []int{1,2,4,8,16,32,64,128}

/*
	Here we make use of range in our for-loop.

	Range returns two things:
	- index
	- value

*/
func main() {
	/*
		Prints out the powers where each
		index is the exponent of 2 and its
		corresponding element is the value
	*/
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i,v)
	}
}
package main

import "fmt"

/*
	1. Default values for types goes as follows:
		0 for numeric types
		false for booleans
		"" for strings
*/

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

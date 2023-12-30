package main

import "fmt"

// Can declare initial values in list format
var i, j int = 1, 2

func main() {

	/*
		1. Can define initial values with different types
		2. Can omit type when initializer is used
	*/

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

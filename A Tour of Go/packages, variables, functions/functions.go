package main

import "fmt"


/*
	Note: types go after variable names
	- this goes before both parameters and functions
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}
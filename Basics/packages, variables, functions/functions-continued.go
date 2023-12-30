package main

import "fmt"

/*
	1. Observe that if we have two consecutive variables with
		the same type, we can omit the type for all variables
		but the last
*/
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

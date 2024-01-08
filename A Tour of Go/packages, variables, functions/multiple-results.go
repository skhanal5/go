package main

import "fmt"

/*
	1. Since we are returning two values, we
		specify types after parameters in parantheses
*/
func swap(x,y string) (string,string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a,b)
}
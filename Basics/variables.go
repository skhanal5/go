package main

import "fmt"

/*
	1. var lets you define function/package level variables
	2. Follows type convention (see bool at the end)
*/

var c, python, java bool 	//package level

func main() {
	var i int 	// function level, default value is 0
	fmt.Println(i,c,python,java)
}
package main

import "fmt"

func main() {
	sum := 0

	/*
		For is the only looping
		structure in Go
	*/

	for i:=0; i<10; i++ {
		//i has block scope
		sum += i
	}	// braces are required

	fmt.Println(sum)
}
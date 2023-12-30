package main

import "fmt"

func main() {
	sum := 2

	/*
		Simplified syntax where
		the init and post statmements
		are gone.

		init -> initializer
		post -> increment/decrement function
	
		Note: semicolon positions
	*/

	for ; sum <10 ; {
		sum += sum
	}	

	fmt.Println(sum)
}
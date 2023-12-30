package main

import "fmt"

func main() {

	/*
		defer waits for the following
		function to return before
		executing the one associated
		with the keyword

		arguments are evaluated immediately
		though...
	*/
	defer fmt.Println("world")

	fmt.Println("hello")
}
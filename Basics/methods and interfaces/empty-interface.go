package main

import "fmt"

func main() {
	
	/*
		interface{} is the empty interface.

		a variable of type empty interface
		can hold values of any type.

		This is good when you want to handle
		situations where you don't know what type
		you are dealing with.

		Example: fmt.Printf
	*/
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
                                                                                                                                                                 
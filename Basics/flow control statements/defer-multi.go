package main

import "fmt"

func main() {
	fmt.Println("counting")

	/*
		Each 'defer' statement
		goes into a stack.

		They are processed in LIFO order
	*/
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

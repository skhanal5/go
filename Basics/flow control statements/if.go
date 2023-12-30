package main

import "fmt"

func main() {

	x := 1

	// Same rule as for where {'s are needed
	if x < 0 {
		fmt.Println("Hello")
	} else {
		fmt.Println("Goodbye")
	}
}

package main

import "fmt"

func main() {

	/*
		syntax to define an array: name [size]type
		Note: the spacing between size and type

		Arrays cannot be resized in Go
	*/
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)	// when printing, it isn't comma separated

	primes := [6]int{2, 3, 5, 7, 11, 13} //defining specific values
	fmt.Println(primes)
}

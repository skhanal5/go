package main

import "fmt"

func main() {
	var i interface{} = "hello"

	/*
		type assertions allow us to access an interface value's underlying
		value.

		The following line is an example of a type assertion
		Assertion syntax: t := i.(T)

		This is basically saying that the value of interface i holds type
		T and that we want to assign the underlying T value
		to t.

		In the line below, we are saying we want the string value
		of the interface

		What happens if i isn't really of type T, we get a panic
	*/
	s := i.(string)
	fmt.Println(s)

	/*
		When we want to test if a value, the type assertion 
		returns the underlying value and a boolean
		that indicates if our assertion passed

		t,ok := i.(T)

		if i holds T, t the value of type T and ok is bool
		else t is zero value of type T, ok is false
	*/
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

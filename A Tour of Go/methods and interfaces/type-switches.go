package main

import "fmt"

func do(i interface{})    {

	/*
		This is a type switch. We can do
		many type assertions in a row using this

		syntax: v := i.(type)
		notice its type and not T

		the cases are all individual types

		If one case passes, then v will have type T in that
		block

		In default case, the variable v is of the same type
		and value as i
	*/
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

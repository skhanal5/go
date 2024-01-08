package main

import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Print("Go runs on: ")

	/*
		Switch is another control flow
		statemnt

		Like if and for, we can initialize
		in the guard

		In Go, it will only run the selected
		case. After it runs that case, it pretty
		much has a implicit "break" at the end,
		meaning subsequent checks don't occur. 

		Switch cases can be non constant and
		values don't need to be integer.
	*/

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}
}
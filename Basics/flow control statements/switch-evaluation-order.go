package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()
	
	/*
		In Go, there is no "falling through"
		behavior where if one case applies
		and it doesn't have a code blocl, it
		doesn't run another case that also applies
		but does have a code block.

	*/
	
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("God knows when")

	}
}
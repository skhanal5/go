package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	/*
		Notice our switch
		statement doesn't have
		a condition. This is the
		same as "true" or in this
		case switch true
	*/
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

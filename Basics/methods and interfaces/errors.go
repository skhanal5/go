package main

import (
	"fmt"
	"time"
)

// custom struct to represent an error
type MyError struct {
	When time.Time
	What string
}

/*
	This implements the error state with error values
	-> error type has one method signature: Error() string

	fmt package looks for error interface when printing error values
*/
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// function that always returns our error struct
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/*
	Functions can return an error value
	
	Calling code should handle errors by testing if
	the error returns nil
*/
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
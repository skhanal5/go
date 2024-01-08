package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() { 
	head := List[int]{next:nil,val: 1}
	fmt.Printf("%d", head.val)
}

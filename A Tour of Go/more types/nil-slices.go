package main

import "fmt"

func main() {
	/*
		We have a nil slice below.
		Length is 0
		Capacity is nil
		
		No underlying array is created
	*/	
	var s []int 
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

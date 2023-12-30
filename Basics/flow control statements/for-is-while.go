package main

import "fmt"

func main() {
	sum := 2

	/*
		We can just remove the ;'s
		
		for's alternative syntax
		is while in other languages
	*/

	for sum <10 {
		sum += sum
	}	

	fmt.Println(sum)
}
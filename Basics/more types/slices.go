package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	/*
		[]type is a "slice"

		Slices are flexible and dynamically sized
		- more common than Arrays in practice

		You form a slice like this: [low bound : high bound]
		where high bound isn't inclusive in the slice
	*/
	var s []int = primes[1:4]
	fmt.Println(s)
}

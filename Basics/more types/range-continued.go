package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	/*
		We can put a _ in either the
		index or value portion that is 
		returned by range. This will
		"ignore" that returned value
	*/
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	/*
		If we want just the index,
		just omit the value portion
	*/
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}

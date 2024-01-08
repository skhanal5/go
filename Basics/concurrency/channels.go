package main

import "fmt"

/*
	Channels allow us to send and receive values using <-
	- they are 'typed conduits'
	- data flows into the direction of the arrow

	Channels MUST be created before use
	syntax: ch := make(chan int)

	Send/Receive operations are blocking untl the other side is
	ready
	- goroutines can synchronize without locks                   
*/

// gets the sum of all numbers in the slice
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
	Here we have three routines.
	main(), go sum 1, go sum 2
	
	go sum 1 will send its sum (the first half) and block until main()
	receives it. it will then terminate

	go sum 2 will send its sum (the second half) and block until main()
	receives it. it will then terminate

	main will print out both sums

*/
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

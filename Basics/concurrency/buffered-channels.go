package main

import "fmt"

/*
	Channels can be buffered
	The second argument to make is the length of the buffer

	When you send to a buffer channel, it'll block only
	if the buffer is full

	Receive blocks when the buffer is empty

	If you try to send too much stuff into a buffered channel,
	you will get deadlock
*/
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

                                                        
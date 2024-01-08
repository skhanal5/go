package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {

		/*
			Select lets a goroutine wait on multiple
			operations

			select blocks until one of its cases can run
			- then it executes it
			- if multiple are ready, it'll pick one at random

		*/
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*
	We first make two UNBUFFERED channels 
	- Recall unbuffered = sender blocks until the receiver gets the value

	Remember main() and go func() are independent and running at the
	same time

	So fibonacci is making values and sending them (via main)
	
	1. When one value is sent, it blocks fibonacci until go func() receives
	and prints it out
	2. Another value is sent, it blocks fibonacci until go func() receives
	and prints it out
	3. continue above step until we have 10 values
	4. It is at this time go func*() passes 0 to fibonacci
	5. switch receives from quit -> terminates function
*/
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

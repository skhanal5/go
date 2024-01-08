package main

import(
	"fmt"
)

func fibonacci(n int, c chan int){
	x,y := 0,1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// senders can close a channel
	// to indicate no more values will be sent
	close(c)
}

/*
	Closing channel's isn't the same as closing a file.
	It isn't necessary really; only in cases like the 
	one you see below it makes sense
*/

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	
	/*
		This loop will run until the channel is closed
	*/
	for i := range c {
		fmt.Println(i)
	}
}
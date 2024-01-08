package main

import (
"fmt"
"time"
)

func say(s string) {
	for i := 0; i<5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
	a goroutine is a lightweight thread managed by Go runtime

	syntax: go [expression]
	Example: go f(x,y,z)

	the above starts a new goroutine that runs f(x,y,z)
	EVALUATION of f,x,y,z happens in the CURRENT goroutine
	EXECUTION happens in the NEW goroutine

	Goroutine's run in the same address space
	- access to shared memory must be synchronized
	- peep sync package

*/
func main() {
	go say("world")
	say("hello")
}
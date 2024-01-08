package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

/*
	Up to here, we have made a type T
	and implemented the method M which
	is defined in our interface.

	Unlike other languages, we haven't declared
	that type T was going to implement the interface,
	we just kind of went along and did it. 

*/
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

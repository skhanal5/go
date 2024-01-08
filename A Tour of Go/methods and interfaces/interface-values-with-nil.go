package main

import "fmt"


/*
	If the value of the interface type
	happens to be nil, the method will be
	called with a nil receiver. Unusually, the
	type of the value is non-nil at this moment.

	In Go: we do not pass execeptions, we write
	methods that handle being called with a nil receiver.
*/
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {

	// nil reciever for M
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T	// *T is nil
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

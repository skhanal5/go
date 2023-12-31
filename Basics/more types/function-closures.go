package main

import "fmt"


/*
	Adder is a closure

	Closures can reference variables
	from outside of its body

	In adder, sum is being referenced in the
	closure

	When a closure references a variable outside
	the body, it binds that variable

*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	/*

		Both pos and neg are two closure's
		where sum is bound to each unique instance.

		We will see in the below example how
		sum will keep track and update state
		through each iteration.
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

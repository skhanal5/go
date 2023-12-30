package main

import "fmt"

const Pi = 3.14

// consts are pretty standard, just can't use :=

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
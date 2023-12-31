package main

import "fmt"

func main() {
	m := make(map[string]int)

	// This is adding a key
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// adding another key
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// deleting a key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	/*
		This will get the value and a bool indicating
		if the key exists

		if the key doesn't exist, the value is 0
	*/
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

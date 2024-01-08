package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v		// create a pointer to the struct
	p.X = 1e9	// we can directly reference the field and assign to it
	fmt.Println(v)
}

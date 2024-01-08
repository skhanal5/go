package main

import "fmt"

type Vertex struct {
	X, Y int
}

// group variable notation (from first section)
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit (this is a literal definition)
	v3 = Vertex{}      // X:0 and Y:0 (default values)
	p  = &Vertex{1, 2} // has type *Vertex 
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

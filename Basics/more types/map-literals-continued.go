package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/*
	Since our top-level type of Vertex
	is just a type name, we can omit it
	in our literal definition. See
	example below
*/
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}


/*
	maps map keys to values
	syntax: map[Type of Key]Type of Value

	a nil map has no key and keys can't be added either

	make returns a map of the given type, initialized, and
	ready to use
*/
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
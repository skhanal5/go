package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
	Like before, ScaleFunc needs the
	first argument to be a pointer
	or else it won't compile

	Note: FUNCTIONS are strict with
	argument types!
*/
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	/*
		One interesting feature of Go is that
		even though 2 is a value, it will still
		invoke the Scale METHOD with a pointer
		receiver.

		This is intentional and the way to 
		interpret this behind the scenes is like
		(&v).Scale(2)
	*/
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

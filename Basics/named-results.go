package main

import "fmt"


/*
	1. You can name your return variables (i.e, x,y)
	2. If you do a return w/o args, it will return named returns
		- recommended only doing this for short functions
*/

func split(sum int) (x,y int) {
	x = sum * 4/9 //integer divison
	y = sum-x
	return
}

func main() {
	fmt.Println(split(17))
}
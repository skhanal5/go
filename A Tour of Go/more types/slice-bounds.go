package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13} //init slice literal

	s = s[1:4]	//get a slice of a slice..
	fmt.Println(s) //[3,5,7]

	s = s[:2]	//get everything up to 2 (not including 2 though)
	fmt.Println(s) //[3,5]

	s = s[1:]	//get everything from 1 (including 1)
	fmt.Println(s) //[5]
}

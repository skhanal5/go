package main

import (
	"fmt"
	"math"
)

/*
	1. Exported names begin with capital letters
	2. You can only import a package and refer to its exported names,
	if the name isn't exported -> it isn't accessible publicly
*/

func main() {
	// fmt.Println(math.pi) pi isn't exported
	fmt.Println(math.Pi)
}

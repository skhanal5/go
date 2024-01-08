package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1,ok1 := <- ch1
		v2 := <- ch2
		if v1 != v2 {
			return false
		}

		if ok1 == false{
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i<10; i++ {
		fmt.Printf("%d\n", <- ch)
	}
	fmt.Printf("%t\n", Same(tree.New(1), tree.New(2)))
}
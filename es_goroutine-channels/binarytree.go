package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	l, r := t.Left, t.Right
	if l != nil {
		Walk(l, ch)
	}
	ch <- t.Value
	if r != nil {
		Walk(r, ch)
	}
	//close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	select {
	case x := <-ch1:
		if x != <-ch2 {
			return false
		}
	case y := <-ch2:
		if y != <-ch1 {
			return false
		}
	}
	return true
}

func main() {
	// First Test
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// Second Test
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

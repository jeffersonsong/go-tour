package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	doWalk(t, ch)
	close(ch)
}

func doWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	doWalk(t.Left, ch)
	ch <- t.Value
	doWalk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v2 != v1 {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func main() {
	result := Same(tree.New(1), tree.New(1))
	fmt.Printf("is same? %v", result)
}

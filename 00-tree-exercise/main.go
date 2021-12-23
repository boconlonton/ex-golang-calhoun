package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walkRecursive func(t *tree.Tree)
	walkRecursive = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walkRecursive(t.Left)
		ch <- t.Value
		walkRecursive(t.Right)
	}
	walkRecursive(t)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	c := make(chan int)

	go Walk(tree.New(1), c)
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}

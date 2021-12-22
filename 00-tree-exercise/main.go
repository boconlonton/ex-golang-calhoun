package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

func main() {
	c := make(chan int)

	go Walk(tree.New(1), c)
	for {
		select {
		case v, ok := <-c:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}

package main

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	var res []int
	c := make(chan int)
	go Walk(tree.New(1), c)

	// Assert the result
	for v := range c {
		res = append(res, v)
	}
	for i, v := range res {
		if i == 0 {
			if v > res[i+1] {
				t.Error()
			}
		} else {
			if v < res[i-1] {
				t.Error()
			}
		}
	}
}

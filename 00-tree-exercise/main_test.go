package main

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	c := make(chan int)
	go Walk(tree.New(1), c)

	// Assert the result
	tmp1, ok1 := <-c
	if !ok1 {
		return
	}
	for {
		tmp2, ok2 := <-c
		if !ok2 {
			break
		}
		if tmp2 < tmp1 {
			t.Error()
		}
	}
}

func TestSame(t *testing.T) {
	r := Same(tree.New(1), tree.New(2))
	if r {
		t.Error()
	}
	r = Same(tree.New(1), tree.New(1))
	if !r {
		t.Error()
	}
	r = Same(tree.New(2), tree.New(1))
	if r {
		t.Error()
	}
}

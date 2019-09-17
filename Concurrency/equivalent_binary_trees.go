package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(k) {
		t = insert(t, (1 + v) * 1)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree {nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + "<-"
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += "->" + t.Right.String()
	}
	return "(" + s + ")"
}



func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	fmt.Println(t.Value)
	if t.Right !=nil {
		Walk(t.Right, ch)
	}
	ch <- 0
}

func Same(t1, t2 *Tree) bool {
	same := true
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)


	for i := range ch1 {
		j := <-ch2
		if i == 0 && j == 0 {
			break
		}
		if (i ==0 && j != 0) || (i != 0 && j == 0) || (i != j) {
			same = false
			break
		}
	}
	return same
}

func main() {
	t1 := New(2)
	t2 := New(2)
	fmt.Println(t1.String())
	fmt.Println(t2.String())
	same := Same(t1, t2)
	fmt.Println(same)

	ch := make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch)
	for i := range ch {
		count := 0
		if i == 0 {
			count++
			if count == 1 {
				break
			}
		}

		fmt.Println("at for loop: ", i)
	}
}
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

// part interface-values.go

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I = T{"hello"}
	i.M()

	var i1 I

	i1 = &T{"Hello"}
	describe(i1)
	i1.M()

	i1 = F(math.Pi)
	describe(i1)
	i1.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

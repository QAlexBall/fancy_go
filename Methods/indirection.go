package main

import (
	"fmt"
)

type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex1, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main() {
	v := Vertex1{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex1{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
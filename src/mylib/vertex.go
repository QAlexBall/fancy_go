package mylib

import (
	"fmt"
	"math"
)

type MyVertex struct {
	X, Y float64
}

func (vertex MyVertex)Show() float64 {
	fmt.Println(vertex.X)
	fmt.Println(vertex.Y)
	return 0
}

func (vertex MyVertex) Abs() float64 {
	return math.Sqrt(vertex.X*vertex.X + vertex.Y*vertex.Y)
}

func (vertex *MyVertex) Scale(f float64) {
	vertex.X = vertex.X * f
	vertex.Y = vertex.Y * f
}

//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
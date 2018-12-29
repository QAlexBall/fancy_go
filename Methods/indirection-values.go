package main

import (
	"fmt"
	"math"
	"mylib"
)

func main() {
	v := mylib.MyVertex{X: 3, Y: 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &mylib.MyVertex{X: 4, Y: 3}
	fmt.Println(p.Abs())	// 解析成(*p).Abs()
	fmt.Println(AbsFunc(*p))


	v1 := &mylib.MyVertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v1, v1.Abs())
	v1.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v1, v1.Abs())
}

func AbsFunc(v mylib.MyVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
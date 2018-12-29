package main

import (
	"fmt"
	"math"
	"mylib"
)

/*
 * 接口类型的变量可以保存任何实现了这些方法的值。
 */

type MyFloat float64

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := mylib.MyVertex{X: 3, Y: 4}

	a = f
	a = &v

	// a = v
	fmt.Println(a.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
package main

import (
	"fmt"
	"mylib"
)

func main() {
	mylib.Test()
	vertex := mylib.MyVertex{X: 3, Y: 4}
	fmt.Println(vertex.X)
	vertex.Show()
}
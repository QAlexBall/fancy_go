package main

import (
	"fmt"
	"mylib"
)

func main() {
	vertex := mylib.MyVertex{X: 3, Y: 4}
	fmt.Println(vertex.X)
	vertex.Show()
}
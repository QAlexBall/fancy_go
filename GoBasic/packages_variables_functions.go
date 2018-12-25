package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func myPrint(x int, y string) (int, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func shortVariableDeclaration() {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

var (
	ToBe          = false
	MaxInt uint64 = 1 << 64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println("hello, world!")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(12, 25))

	a, b := swap("hello", "world")
	a1, b1 := myPrint(3, b)
	fmt.Println(a, b)
	fmt.Println(a1, b1)
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
	shortVariableDeclaration()

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	{
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
	}

	{
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x * x + y * y))
		var z uint = uint(f)
		fmt.Println(x, y, z)
	}

	{
		var x, y = 3, 4
		var f = math.Sqrt(float64(x * x + y * y))
		var z = uint(f)
		fmt.Println(x, y, z)
	}

	{
		v := 32.1
		fmt.Printf("v is of type %T\n", v)
	}

	{
		const Pi = 3.14
		{
			const World = "世界"
			fmt.Println("Hello", World)
			fmt.Println("Happy", Pi, "Day")

			const Truth = true
			fmt.Println("Go rules?", Truth)
		}
	}

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
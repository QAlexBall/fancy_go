package main

import (
	"fmt"
	"math"
	"strings"
)

// struct
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}	// has type Vertex
	v2 = Vertex{X: 1}		// Y:0 is implicit
	v3 = Vertex{}			// X:0 and Y:0
	p = &Vertex{1, 2}	// has type *Vertex
)

func point() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

var m1 = map[string]Vertex1 {
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	point()
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	{
		v := Vertex{1, 2}
		p := &v
		p.X = 1e9
		fmt.Println(v)
	}

	fmt.Println(v1, p, v2, v3)

	{
		var a [2]string
		a[0] = "Hello"
		a[1] = "World"
		fmt.Println(a[0], a[1])
		fmt.Println(a)

		primes := [6]int{2, 3, 5, 7, 11, 13}
		fmt.Println(primes)
	}

	{
		primes := [6]int{2, 3, 5, 7, 11, 13}
		var s = primes[1:4]
		fmt.Println(s)
	}

	{
		names := [4]string {
			"John",
			"Paul",
			"George",
			"Ringo",
		}
		fmt.Println(names)

		a := names[0:2]
		b := names[1:3]
		fmt.Println(a, b)

		b[0] = "XXX"
		fmt.Println(a, b)
		fmt.Println(names)
	}

	{
		q := []int{2, 3, 5, 7, 11, 13}
		fmt.Println(q)

		r := []bool{true, false, true, true, false, true}
		fmt.Println(r)

		s := []struct {
			i int
			b bool
		} {
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println(s)
	}

	{
		s := []int{2, 3, 5, 7, 11, 13}

		s = s[:2]
		fmt.Println(s)

		s = s[1:]
		fmt.Println(s)
	}

	{
		s := []int{2, 3, 5, 7, 11, 13}
		printSlice(s)

		s = s[:0]
		printSlice(s)

		s = s[:4]
		printSlice(s)

		s = s[2:]
		printSlice(s)
	}

	{
		var s []int
		fmt.Println(s, len(s), cap(s))
		if s == nil {
			fmt.Println("nil")
		}
	}

	{
		a := make([]int, 5)
		printSlice1("a", a)

		b := make([]int, 0, 5)
		printSlice1("b", b)

		c := b[:2]
		printSlice1("c", c)

		d := c[2:5]
		printSlice1("d", d)
	}

	{
		board := [][]string{
			{"_", "_", "_"},
			{"_", "_", "_"},
			{"_", "_", "_"},
		}

		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"
		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
	}

	{
		var s []int
		printSlice2(s)

		s = append(s, 0)
		printSlice2(s)

		s = append(s, 1)
		printSlice2(s)

		s = append(s, 2, 3, 4)
		printSlice2(s)
	}

	{
		for i, v := range pow2 {
			fmt.Printf("2 ** %d = %d\n", i, v)
		}
	}

	{
		pow := make([]int, 10)
		for i := range pow {
			pow[i] = 1 << uint(i)
		}
		for _, value := range pow {
			fmt.Printf("%d\n", value)
		}
	}

	{
		m = make(map[string]Vertex1)
		m["Bell Labs"] = Vertex1{
			40.68433, -74.39967,
		}
		fmt.Println(m)
	}
	fmt.Println(m1)

	{
		m := make(map[string]int)
		m["Answer"] = 42
		fmt.Println("The value:", m["Answer"])

		m["Answer"] = 48
		fmt.Println("The value:", m["Answer"])

		delete(m, "Answer")
		fmt.Println("The value:", m["Answer"])

		v, ok := m["Answer"]
		fmt.Println("The value:", v, "Present?", ok)
	}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	{
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(
				pos(i),
				neg(-2*i),
				)
		}
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func printSlice2(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}

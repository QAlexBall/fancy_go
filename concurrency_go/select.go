package main

import (
	"fmt"
	"time"
)

// select 专门用于处理通道

func selectTest1() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func selectTest2() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func selectTest3() {
	done := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

	workCounter := 0
	loop :
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		//time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

func main() {
	//selectTest1()
	//selectTest2()
	selectTest3()
}
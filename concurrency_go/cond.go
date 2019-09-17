package main

import (
	"fmt"
	"sync"
	"time"
)

// cond 实现了一个条件变量用于等待或宣布事件发生时goroutine的交汇点

func removeTest()  {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("'Adding to queue")
		queue = append(queue, struct {}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func buttonTest()  {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var tempwg sync.WaitGroup
		tempwg.Add(1)
		go func() {
			tempwg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		tempwg.Wait()
	}

	var wg sync.WaitGroup
	wg.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		wg.Done()
	})

	button.Clicked.Broadcast()

	wg.Wait()
}


func main() {
	buttonTest()
}
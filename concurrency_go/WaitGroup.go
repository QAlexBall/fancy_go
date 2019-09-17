package main

import (
	"fmt"
	"sync"
	"time"
)

//你可以把WaitGroup视作一个安全的并发计数器：调用Add增加计数，调用Done减少计数。
// 调用Wait会阻塞并等待至计数器归零。

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
	fmt.Println()


	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg1 sync.WaitGroup
	wg1.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg1, i + 1)
	}
	wg1.Wait()
}


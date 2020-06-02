package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create one buffered channel
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	// start all goroutines
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add a group of work
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d: Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker %d: Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		fmt.Printf("Worker: %d: Completed %s\n", worker, task)
	}
}

package main

import (
	"bytes"
	"fmt"
	"sync"
)

func constraintTest1 () {
	data := make([]int, 4)
	data = append(data, 1, 2, 3, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func constraintTest2 () {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // return 一个只能被单向读取的通道
		  							 // chan<- int 一个只能单向发送的通道
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}

func constraintTest3 () {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[:3])
	wg.Wait()
}

func main() {
	constraintTest3()
}
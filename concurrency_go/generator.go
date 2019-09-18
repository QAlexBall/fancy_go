package main

import (
	"fmt"
)

func repeat (done <-chan interface{}, values ...interface{}) <-chan interface{} {

	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func take (done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func toString (done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}


func testGenerator1 () {
	done := make(chan interface{})
	defer close(done)


	for num := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Printf("%v ", num)
	}
}

func testGenerator2 () {
	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I ", "am ", "22", " years old."), 8)) {
		message += token
	}
	fmt.Printf("[message] -> %s...", message)
}

func main() {
	testGenerator2()
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func repeatFn (done <-chan interface{}, fn func() interface{}) <-chan interface{} {

	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toInt(done <-chan interface{}, repeatStream <-chan interface{}) <-chan int{
	randIntStream := make(chan int)
	go func() {
		defer close(randIntStream)
	}()
	return randIntStream
}

func primeFinder(done <-chan interface{}, randIntStream <-chan int) <-chan interface{} {

}

func main() {
	rand := func() interface{} { return rand.Intn(500000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes: ")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}

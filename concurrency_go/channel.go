package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

/*
 * Channel, 即通道，衍生自CSP并发模型，是Go的并发原语，在Go语言中具有及其重要的地位。
 * dataStream = make(chan interface{}) 这里使用内置函数make实力化通道
 */

func test1() {
	stringStream := make(chan string)
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		close(stringStream)
		stringStream <- "Hello channels"
	}()
	wg.Wait()
	salutation, ok := <-stringStream // ok 是读取操作的一个标识，
									 // 用于指示读取的通道是
									 // 由过程中其他位置的写入生成的值，
									 // 还是由已关闭通道生成的默认值。
	fmt.Printf("(%v): %v", ok, salutation)
}

func test2()  {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Println(integer)
	}
}

func test3()  {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i:= 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func test4() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Seding: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
/*
 * 当一个goroutine拥有一个通道时应该：
 * 1。 初始化
 * 2。 执行写入操作，将所有权交给另一个goroutine
 * 3。 关闭该通道
 * 4。 将此前列出的三个事情封装在一个列表中，通过订阅通道将其公开
 */

func test5() {
	chanOwner := func() <-chan int {
		resultSteam := make(chan int, 5)
		go func() {
			defer close(resultSteam)
			for i:= 0; i <= 5; i++ {
				resultSteam <- i
			}
		}()
		return resultSteam
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func main() {
	//test2()
	//test3()
	//test4()
	test5()
}
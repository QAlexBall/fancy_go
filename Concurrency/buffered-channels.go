package main

import "fmt"

// 信道是可以带缓冲的
func main() {
	ch := make(chan int, 2)	// 将缓冲长度作为第二个参数提供给make来初始化一个带缓冲的信道
	ch <- 1
	ch <- 2
	// ch <- 3		fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

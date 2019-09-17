package main

import (
	"fmt"
	"sync"
)

// Pool是对象池模式的并发安全实现。
func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create new instance.")
			return struct {}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}
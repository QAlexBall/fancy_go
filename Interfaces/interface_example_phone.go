package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct { }

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I called you!")
}

type IPhone struct { }

func (iphone IPhone) call() {
	fmt.Println("I am IPhone, I called you!")
}

func main() {
	var phone Phone
	var iphone Phone

	phone = new(NokiaPhone)
	phone.call()

	iphone = new(IPhone)
	iphone.call()
}

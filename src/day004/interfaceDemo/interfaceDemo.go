package main

import "fmt"

/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，
任何其他类型只要实现了这些方法就是实现了这个接口。
*/

type Phone interface {
	call()
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am IPhone, I can call you")
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am NokiaPhone, I can call you")
}

func main() {
	var phone Phone

	phone = new(IPhone)
	phone.call()

	phone = new(NokiaPhone)
	phone.call()
}

package main

import "fmt"

//声明一个函数类型
type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x:%d\n", x)
	return x
}

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x:%d\n", x)
		return x
	})
}

//我是回调，x:1
//我是回调，x:2

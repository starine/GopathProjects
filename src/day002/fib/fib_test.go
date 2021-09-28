package fib

import (
	"fmt"
	"testing"
)

//实现Fibonacci数列，1，1，2，3，5，8，13
func TestFibList(t *testing.T){
	//var a int = 1
	//var b int = 1

	//var(
	//	a int = 1
	//	b int = 1
	//)

	a := 1
	b := 1

	fmt.Print(a)
	for i := 0; i < 10; i++ {
		fmt.Print(" ",b)
		//tmp := a
		//a = b
		//b = tmp + a

		a, b = b, a+b
		//b, a = a+b, b //没有差别
	}
	t.Log("Fibonacci")
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
	t.Log(a,b)
}
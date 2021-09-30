package main

import "fmt"

//阶乘
func Factorial(n uint64) uint64 {
	if n > 0 {
		result := n * Factorial(n-1)
		return result
	}
	return 1
}

//斐波那契数列
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func main() {
	//阶乘
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", Fibonacci(i))
	}
}

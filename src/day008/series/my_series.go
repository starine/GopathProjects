package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

//小写开头的方法在包外不可见
func square(n int) int {
	return n * n
}

var LessThanOneError = errors.New("n should be not less than 1")
var LargeThanOneError = errors.New("n should be not large than 100")

func GetFibonacciSerie(n int) ([]int, error) {
	if n < 1 {
		return nil, LessThanOneError
	}
	if n > 100 {
		return nil, LargeThanOneError
	}
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret[:n], nil
}

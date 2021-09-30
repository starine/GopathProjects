package main

import (
	"errors"
	"fmt"
	"math"
)

/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：
type error interface {
    Error() string
}
我们可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
*/

type DivideError struct { // 定义 DivideError 结构体
	dividee int
	divider int
}

func (de *DivideError) Error() string { //DivideError 结构体 实现error接口
	strFormat := "Cannot proceed, the divider is zero\n\tdividee: %d\n\tdivider:0"
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{dividee: varDividee, divider: varDivider}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func mySqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), errors.New("")
}

func main() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is:", errorMsg)
	}
	if ret, err := mySqrt(-4.0); ret == 0 {
		fmt.Println("errorMsg is:", err)
	}
}

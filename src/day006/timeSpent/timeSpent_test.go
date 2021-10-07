package timeSpent_test

import (
	"fmt"
	"testing"
	"time"
)

//计算函数运行时长

// 输入是函数类型，返回也是函数类型
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		// 输出函数运行时间
		fmt.Println("time spent :", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op * op * op
}

func TestFunc(t *testing.T) {
	fnSF := timeSpent(slowFunc) //传入函数对象，返回函数对象
	t.Log(fnSF(10))
}

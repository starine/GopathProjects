package defer_test

import "testing"

func TestDefer(t *testing.T) {
	defer func() { // 匿名函数，也可以传其他函数
		t.Log("clear resources")
	}()
	t.Log("Started")
	panic("fatal error") // defer 仍然会执行
}

/*
类似于try{}except{}final{}的结构
final会被执行
panic相当于raise error了
如果有panic，那么执行了panic抛出了错误，defer仍然会执行，但是panic之后的语句，不会再被执行
可用defer安全的释放资源释放锁～
*/

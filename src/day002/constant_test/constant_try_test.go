package constant_test

import "testing"

const(
	Monday = 1 + iota
	Tuesday
	Wednesday
)
const(
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //二进制0111
	t.Log(a&Readable, a&Writable, a&Executable) //(&)按位与运算，两个都为1，结果才为1，相当于串联
	//(|)按位或运算，两个都为0，结果才为0 ，相当于并联
	//(^)按位异或运算，相同为0，不同为1
	//(~)非运算，取反
}
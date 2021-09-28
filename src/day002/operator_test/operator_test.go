package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 4, 3}

	t.Log(a == b)  	//true
	t.Log(a == d)	//false
	//t.Log(a == c) //invalid operation: a == c (mismatched types [4]int and [5]int)
}

/*
&^ 按位置零,
1 &^ 0  = 1
1 &^ 1  = 0
0 &^ 1  = 0
0 &^ 0  = 0
右边为1，结果都为；
右边为0，结果与左边相同
 */
const(
	Readable = 1 << iota
	Writable
	Executable
)
func TestBitClear(t *testing.T) {
	a := 7 //二进制0111
	a = a &^ Readable //按位清零
	a = a &^ Writable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable) //
}
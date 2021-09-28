package type_test

import "testing"

type MyInt int64 //别名

//go不支持隐式类型的转换，只能显示转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 3
	//b = a //不行，因为go不支持隐式类型的转换
	b = int64(a) //显示的类型转换是可以的


	var c MyInt
	//c = b //cannot use b (type int64) as type MyInt in assignment，对于别名的隐式类型转换也是不行的
	c = MyInt(b)
	t.Log(a, b, c)
}

//go不支持指针的运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr +1 //invalid operation: aPtr + 1 (mismatched types *int and int) 不支持指针运算
	t.Log(a,aPtr)
	t.Logf("%T %T", a, aPtr)
}

//string 未初始化默认是空字符串“”，而不是nil
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	//if s == ""{
	//
	//}

}
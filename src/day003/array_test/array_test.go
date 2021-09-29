package array_test

import "testing"

//数组是一块连续的存储空间，数组是具有相同唯一类型的一组已编号且长度固定的数据项序列
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr1[2] = 6
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2, arr)
}

//数组的遍历方式
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	//通用方式
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	//range 返回数组的索引和其存放的值
	//for idx, e := range arr3 {
	//	t.Log(idx, e)
	//}
	for _, e := range arr3 {
		t.Log(e)
	}
}

//s := arr[startIndex:endIndex]
func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	arr4_sec := arr4[3:]
	t.Log(arr4_sec)
}

func TestArrayCompare(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 8, 4}
	//d := [...]int {1, 2, 3, 4, 5}
	t.Log(a == b) //array_test.go:41: true
	t.Log(b == c) //array_test.go:44: false
	//t.Log(b == d)//invalid operation: b == c (mismatched types [4]int and [5]int)

	//相同长度数组是可以比较的，不同长度数组无法比较
}

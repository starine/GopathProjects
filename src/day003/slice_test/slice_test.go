package slice_test

import "testing"

//Go数组长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活的、功能强悍的内置型切片("动态数组")
//与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
func TestSliceInit(t *testing.T) {
	//slice的声明和数组非常像，不同的是slice声明时[]中没有值
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3])//index out of range [3] with length 3
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) //append后若length超过原来的capcity，那么capacity的大小变成原来capacity的两倍，成倍增长
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // slice_test.go:40: [Apr May Jun] 3 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) //slice_test.go:42: [Jul Aug Sep] 3 7
	//Q2和Summer是year的一个切片，它们是共享存储空间的，所以Q2的capacity是从切片的初始地址一直到year的capacity末尾

	summer[0] = "unknow"
	t.Log(Q2) //slice_test.go:46: [Apr May unknow]
	//在summer中修改值，其实也修改了year和Q2的中的值，因为他们共享一块连续的存储空间
	//slice是一个结构体，其中data存放的是一块连续存储空间的首地址，len表示切片的长度，cap表示切片的容量

	t.Logf("year[5]的值为%s, Q2[2]的值为%s, summer[0]的值为%s\n", year[5], Q2[2], summer[0])
	t.Logf("year[5]的地址为%d, Q2[2]的地址为%d, summer[0]的地址为%d\n", &year[5], &Q2[2], &summer[0])
	/*运行结果
	slice_test.go:50: year[5]的值为unknow, Q2[2]的值为unknow, summer[0]的值为unknow
	slice_test.go:51: year[5]的地址为824634916944, Q2[2]的地址为824634916944, summer[0]的地址为824634916944
	*/

	//year = append(year,"loop")//append后若length超过原来的capcity，那么capacity的大小变成原来capacity的两倍
	//t.Log(len(year),cap(year))//slice_test.go:42: 13 24

}

func TestSliceComparing(t *testing.T) {
	//a := []int {1, 2, 3, 4}
	//b := []int {1, 2, 3, 4}

	//if a == b {  //invalid operation: a == b (slice can only be compared to nil)
	//	t.Log(a == b)
	//}
	//slice是不可以比较的，它只能和 nil 比较

	//s1 := a[1:3]
	//s2 := a[1:3]
	//t.Log(s1 == s2) //invalid operation: s1 == s2 (slice can only be compared to nil)
}

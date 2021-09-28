package select__test

import "testing"

func TestSelect(t *testing.T) {
	var c1, c2, c3 chan int
	var i1, i2 int

	//通道（channel）是用来传递数据的一个数据结构，操作符<-用于指定通道的方向
	select {
		case i1 = <- c1 : //从 c1 接收数据，并把值赋给 i1,向chan传入数据
			t.Log("received", i1, "from c1")
		case c2 <- i2:  //把 i2 发送到通道 c2，向chan传入数据
			t.Log("sent", i2, "to c2")
		case i3, ok := <-c3:
			if ok {
				t.Log("received", i3, "from c3")
			}else{
				t.Log("c3 is closed")
			}
		default:
			t.Log("no communication")
	}
}

package condition__test

import "testing"

// if condition {}  condition必须为布尔值，0或1不行
func TestIfMultiSec(t *testing.T) {
	//if v, err := someFun(); err == nil {
	//	t.Log("err,nil")
	//}else {
	//	t.Log("err")
	//}
	if a := 1 == 1; a{
		t.Log("1==1")
		t.Log(a)//true
	}

}

//case中不用添加break，默认就有
func TestSwitchMultiCase(t *testing.T) {
	for i := 0;i<5;i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log("is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++{
		switch {
		case i % 2 == 0:
			t.Log("even")
		case i % 2 ==1:
			t.Log("odd")
		default:
			t.Log("unknow")
		}
	}
}

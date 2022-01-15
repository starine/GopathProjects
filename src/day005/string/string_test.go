package string_test

import (
	"testing"
	"unsafe"
)

/*
string 是数据类型， 不是引用或指针类型
string 是只读的byte slice， len函数可以得到它所包含的byte数
string 的byte数组可以存放任何数据

Unicode是一种字符集（字符编码 code point）
UTF8是Unicode的存储实现（转换为字节序列的规则）

*/

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为”“
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' //Cannot assign to s[1],string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	s2 := "\xE4\xBA\xB5\xFF"
	t.Log(s)
	t.Log(len(s))
	t.Log(s2)
	t.Log(len(s2))
	s3 := "中"
	t.Log(len(s3)) //3 len()输出的是byte数量而不是字符数量

	c := []rune(s3) //可以得到字符串的unicode
	t.Log(len(c))   //1
	t.Log("rune size", unsafe.Sizeof(c[0]))
	t.Logf("'中'的 unicode 是 %x", c[0])
	t.Logf("'中'的 UTF8 是 %x", s3)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for i, c := range s { //range每次取的是rune而不是byte
		t.Logf("%d %[2]c %[2]x", i, c)
	}

	for i := 0; i < len(s); i++ { //string 是只读的slice，[]byte
		t.Logf("%d %[2]c %[2]x", i, s[i]) //这里每个汉字占3byte，所以len（s）=21
	}

	c := []rune(s) //转换成rune后，得到每个汉字的unicode
	for i := 0; i < len(c); i++ {
		t.Logf("%d %[2]c %[2]x", i, c[i])
	}
}

package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A, B, C"
	//拆分
	parts := strings.Split(s, ", ")
	for _, part := range parts {
		t.Log(part)
	}
	//连接
	t.Log(strings.Join(parts, "-"))
}

//字符串和整数转换
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str:", s)
	if i, err := strconv.Atoi("5"); err == nil {
		t.Log(10 + i)
	}

}

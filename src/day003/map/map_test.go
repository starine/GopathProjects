package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 4, 3: 8}
	t.Log(m1[2])
	t.Logf("len m1:%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2[4]的值:%d", m2[4])
	t.Logf("m2[2]的值:%d", m2[2])
	t.Logf("len m2:%d", len(m2))
	t.Log(m2)

	m3 := make(map[int]int, 10)
	t.Logf("len m3:%d", len(m3))
}

//在访问的key不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
func TestMapKeyIsExisting(t *testing.T) {
	m4 := map[int]int{}
	m4[2] = 0
	t.Logf("m4[1]的值:%d", m4[1]) //map_test.go:25: m4[1]的值:0

	//判断map中是否存在key，用 v,ok ：= map[key]，ok为true表示存在
	if v, ok := m4[1]; ok {
		t.Logf("m4[1]的值:%d", v)
	} else {
		t.Log("m4 中不存在key为1的键值对") //map_test.go:31: m4 中不存在key为1的键值对
	}

	if v, ok := m4[2]; ok {
		t.Logf("m4[2]的值:%d", v) //map_test.go:35: m4[2]的值:0
	} else {
		t.Log("m4 中不存在key为2的键值对")
	}

}

//map遍历for k, v := range m5{...}
func TestMapTravel(t *testing.T) {
	m5 := map[int]int{1: 3, 2: 9, 3: 27, 4: 81}
	for k, v := range m5 {
		t.Logf("key:%d, value:%d", k, v)
	}
}

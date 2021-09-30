package map_test

import (
	"testing"
)

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，
这是因为 Map 是使用 hash 表来实现的。
*/
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

//第二种遍历方式
func TestMapTravel2(t *testing.T) {
	m5 := map[int]int{1: 3, 2: 9, 3: 27, 4: 81}
	for k := range m5 {
		t.Logf("key:%d, value:%d", k, m5[k])
	}
}

//delete（）函数用于删除集合的元素， 参数为 map 和其对应的 key。
func TestDeleteMap(t *testing.T) {
	//创建map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	t.Log("原始地图")

	//打印地图
	for country := range countryCapitalMap {
		t.Log(country, "首都是", countryCapitalMap[country])
	}
	//删除元素
	delete(countryCapitalMap, "France")
	t.Log("法国条目被删除")
	t.Log("删除后的地图")
	//打印地图
	for country := range countryCapitalMap {
		t.Log(country, "首都是", countryCapitalMap[country])
	}
}

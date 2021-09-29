package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(pa int) int{}
	m[1] = func(pa int) int {
		return pa
	}
	m[2] = func(pa int) int {
		return pa * pa
	}
	m[3] = func(pa int) int {
		return pa * pa * pa
	}
	t.Log(m[1](2), m[2]((2)), m[3](2)) //map_ext_test.go:16: 2 4 8
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 3)
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

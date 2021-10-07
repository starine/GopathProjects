package variableParam

import "testing"

func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

func TestVariableParamFunc(t *testing.T) {
	t.Log(sum(1, 2, 3))
	t.Log(sum(1, 2, 3, 4))
}

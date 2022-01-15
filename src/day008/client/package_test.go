package client

import "testing"
import "day008/series"

func TestPackage(t *testing.T) {
	n := 1
	if v, e := series.GetFibonacciSerie(n); e != nil {
		t.Error(e)
	} else {
		t.Log(v)
	}

	//t.Log(series.square(5)) //cannot refer to unexported name series.square,小写开头的方法在包外不可见
}

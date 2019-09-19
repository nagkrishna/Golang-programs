package gotest

import "testing"

func Test_Div_1(t *testing.T) {
	if i, e := Div(6, 2); i != 3 || e != nil {
		t.Error("division function tests do not pass")
	} else {
		t.Log("first test passed")
	}
}

func Test_Div_2(t *testing.T) {
	t.Error("just does not pass")
}

package test

import "testing"

func hi() string {
	return "hi"
}

func TestRequest(t *testing.T) {
	//input := ""
	expectedOut := "hi"
	actualOut := hi()

	if actualOut != expectedOut {
		t.Errorf("expected output to be %s, but got %s", expectedOut, actualOut)
	} else {
		t.Logf("testcase passed for 1")
	}
}
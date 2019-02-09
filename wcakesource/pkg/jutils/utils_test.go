package jutils

import (
	"testing"
)

func TestCombineString(t *testing.T) {
	ret := CombineString("aaa", "bbb")
	if ret != "aaabbb" {
		t.Fatal(ret)
	} else {
		t.Log(ret)
		t.Log("Done")
	}
}

package leetcode

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	var ret int
	var s string

	s = ""
	ret = 0
	if ret != lengthOfLastWord(s) {
		t.Fatalf("case fails %v\n", ret)
	}

	s = "Hello World"
	ret = 5
	if ret != lengthOfLastWord(s) {
		t.Fatalf("case fails %v\n", ret)
	}
}

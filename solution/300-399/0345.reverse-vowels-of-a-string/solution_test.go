package leetcode

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	var ret string
	var s string

	s = "hello"
	ret = "holle"
	if ret != reverseVowels(s) {
		t.Fatalf("case fails %v\n", ret)
	}

}

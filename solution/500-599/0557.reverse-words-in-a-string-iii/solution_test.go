package leetcode

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	var s string
	var ret string
	s = "Let's take LeetCode contest"
	ret = "s'teL ekat edoCteeL tsetnoc"
	if ret != reverseWords(s) {
		t.Fatalf("case fails %v\n", ret)
	}
}

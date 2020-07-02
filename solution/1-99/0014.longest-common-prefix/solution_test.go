package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var s []string
	var ret string
	s = []string{"flower", "flow", "flight"}
	ret = "fl"
	if ret != longestCommonPrefix(s) {
		t.Fatalf("case fails %v\n", ret)
	}

	s = []string{}
	ret = ""
	if ret != longestCommonPrefix(s) {
		t.Fatalf("case fails %v\n", ret)
	}
}

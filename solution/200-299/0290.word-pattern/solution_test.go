package leetcode

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	var ret bool
	var str string
	var pattern string

	str = "dog cat cat dog"
	pattern = "abba"
	ret = true
	if ret != wordPattern(pattern, str) {
		t.Fatalf("case fails %v\n", ret)
	}

	str = "dog cat cat dog"
	pattern = "aaaa"
	ret = false
	if ret != wordPattern(pattern, str) {
		t.Fatalf("case fails %v\n", ret)
	}
}

package leetcode

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	var haystack string
	var needle string
	var ret int

	haystack = "aaa"
	needle = ""
	ret = 0
	if ret != strStr(haystack, needle) {
		t.Fatalf("case fails: %v\n", ret)
	}

	haystack = "aaa"
	needle = "a"
	ret = 0
	if ret != strStr(haystack, needle) {
		t.Fatalf("case fails: %v\n", ret)
	}

	haystack = "hello"
	needle = "ll"
	ret = 2
	if ret != strStr(haystack, needle) {
		t.Fatalf("case fails: %v\n", ret)
	}

	haystack = "aaaaa"
	needle = "bba"
	ret = -1
	if ret != strStr(haystack, needle) {
		t.Fatalf("case fails: %v\n", ret)
	}

	haystack = "mississippi"
	needle = "issipi"
	ret = -1
	if ret != strStr(haystack, needle) {
		t.Fatalf("case fails: %v\n", ret)
	}
}

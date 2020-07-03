package leetcode

import (
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	var ret bool
	var s string
	var ts string

	ts = "add"
	s = "egg"
	ret = true
	if ret != isIsomorphic(s, ts) {
		t.Fatalf("case fails %v\n", ret)
	}

	ts = "bar"
	s = "foo"
	ret = false
	if ret != isIsomorphic(s, ts) {
		t.Fatalf("case fails %v\n", ret)
	}

	ts = "title"
	s = "paper"
	ret = true
	if ret != isIsomorphic(s, ts) {
		t.Fatalf("case fails %v\n", ret)
	}
}

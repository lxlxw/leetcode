package leetcode

import (
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	var ret []string
	var s string

	s = "AAAAAAAAAAA"
	ret = []string{"AAAAAAAAAA"}
	for i, v := range findRepeatedDnaSequences(s) {
		if ret[i] != v {
			t.Fatalf("case fails %v\n", ret)
		}
	}

	s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	ret = []string{"AAAAACCCCC", "CCCCCAAAAA"}
	for i, v := range findRepeatedDnaSequences(s) {
		if ret[i] != v {
			t.Fatalf("case fails %v\n", ret)
		}
	}

}

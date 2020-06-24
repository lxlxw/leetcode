package leetcode

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	var ret []byte
	var s []byte

	s = []byte("hello")
	ret = []byte("olleh")
	reverseString(s)
	for i, v := range s {
		if v != ret[i] {
			t.Fatalf("case fails %v\n", ret)
		}
	}
}

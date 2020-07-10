package leetcode

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	var num uint32
	var ret int

	num = 00000000000000000000000000001011
	ret = 3

	if ret != hammingWeight(num) {
		t.Fatalf("case fails %v\n", ret)
	}
}

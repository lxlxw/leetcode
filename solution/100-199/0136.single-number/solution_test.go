package leetcode

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	var ret int
	var nums []int
	ret = 1
	nums = []int{2, 2, 1}
	if ret != singleNumber(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 4
	nums = []int{4, 1, 2, 1, 2}
	if ret != singleNumber(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

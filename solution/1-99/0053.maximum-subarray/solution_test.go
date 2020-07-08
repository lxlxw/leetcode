package leetcode

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var nums []int
	var ret int

	ret = 10
	nums = []int{1, 3, 6, -5, 4, -3, 2}
	if ret != maxSubArray(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 0
	nums = []int{}
	if ret != maxSubArray(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

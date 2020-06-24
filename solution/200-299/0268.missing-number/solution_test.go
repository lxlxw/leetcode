package leetcode

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	var ret int
	var nums []int

	ret = 2
	nums = []int{3, 0, 1}
	if ret != missingNumber(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 8
	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	if ret != missingNumber(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

package leetcode

import (
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var ret int
	var nums []int

	nums = []int{1, 1, 0, 1, 1, 1}
	ret = 3
	if ret != findMaxConsecutiveOnes(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

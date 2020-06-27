package leetcode

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	var ret int
	var nums []int

	nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret = 49
	if ret != maxArea(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

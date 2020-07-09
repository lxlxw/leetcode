package leetcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var nums []int
	var ret int

	ret = 8
	nums = []int{1, 2, 3, 4, 8, 2, 9}
	if ret != maxProfit(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 0
	nums = []int{}
	if ret != maxProfit(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

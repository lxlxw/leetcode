package leetcode

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	var ret []int
	var nums []int
	ret = []int{0, 0, 0, 1, 1, 1, 2, 2, 2}
	nums = []int{2, 0, 2, 1, 1, 0, 1, 0, 2}
	sortColors(nums)
	for i, v := range nums {
		if ret[i] != v {
			t.Fatalf("case fails %v\n", ret)
		}
	}
}

package leetcode

import (
	"testing"
)

func TestRotate(t *testing.T) {
	var target int
	var ret []int
	var nums []int

	ret = []int{5, 6, 7, 1, 2, 3, 4}
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	target = 3
	rotate(nums, target)
	for i, v := range nums {
		if v != ret[i] {
			t.Fatalf("case fails: %v\n", ret)
		}
	}
}

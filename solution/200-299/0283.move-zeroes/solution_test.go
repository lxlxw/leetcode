package leetcode

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var nums []int
	var ret []int

	nums = []int{0, 1, 0, 3, 12}
	ret = []int{1, 3, 12, 0, 0}
	moveZeroes(nums)
	for k, v := range nums {
		if v != ret[k] {
			t.Fatalf("case fails: %v\n", ret)
		}
	}
}

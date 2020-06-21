package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	var nums []int
	var ret []int

	nums = []int{1, 2, 3}
	ret = []int{1, 2, 4}
	plus := plusOne(nums)
	for k, v := range ret {
		if plus[k] != v {
			t.Fatalf("case fails: %v\n", plus)
		}
	}
	nums = []int{9, 9, 9}
	ret = []int{1, 0, 0, 0}
	plus = plusOne(nums)
	for k, v := range ret {
		if plus[k] != v {
			t.Fatalf("case fails: %v\n", plus)
		}
	}
}

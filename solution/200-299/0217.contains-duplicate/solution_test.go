package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	var nums []int
	var ret bool

	nums = []int{1, 2, 3, 1}
	ret = true
	if ret != containsDuplicate(nums) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 2, 3, 4}
	ret = false
	if ret != containsDuplicate(nums) {
		t.Fatalf("case fails: %v\n", ret)
	}
}

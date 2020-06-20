package leetcode

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var nums []int
	var target int
	var ret int

	nums = []int{1, 3}
	target = 2
	ret = 1
	if ret != searchInsert(nums, target) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 3, 5, 6}
	target = 5
	ret = 2
	if ret != searchInsert(nums, target) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 3, 5, 6}
	target = 2
	ret = 1
	if ret != searchInsert(nums, target) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 3, 5, 6}
	target = 7
	ret = 4
	if ret != searchInsert(nums, target) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 3, 5, 6}
	target = 0
	ret = 0
	if ret != searchInsert(nums, target) {
		t.Fatalf("case fails: %v\n", ret)
	}
}

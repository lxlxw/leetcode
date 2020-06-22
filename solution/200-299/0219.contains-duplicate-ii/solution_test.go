package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	var nums []int
	var ret bool
	var k int

	nums = []int{1, 2, 3, 1}
	k = 3
	ret = true
	if ret != containsNearbyDuplicate(nums, k) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 0, 1, 1}
	k = 1
	ret = true
	if ret != containsNearbyDuplicate(nums, k) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	ret = false
	if ret != containsNearbyDuplicate(nums, k) {
		t.Fatalf("case fails: %v\n", ret)
	}
}

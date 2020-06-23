package leetcode

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var nums []int
	var ret int

	nums = []int{3, 2, 3}
	ret = 3
	if ret != majorityElement(nums) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	ret = 2
	if ret != majorityElement(nums) {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{}
	ret = 0
	if ret != majorityElement(nums) {
		t.Fatalf("case fails: %v\n", ret)
	}
}

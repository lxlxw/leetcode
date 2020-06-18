package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {

	var ret, val int
	nums1 := []int{}
	val = 0
	ret = 0
	if removeElement(nums1, val) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums2 := []int{3, 2, 2, 3}
	val = 3
	ret = 2
	if removeElement(nums2, val) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums3 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	ret = 5
	if removeElement(nums3, val) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}

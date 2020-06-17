package leetcode

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	var ret int
	nums1 := []int{1, 1, 2}
	ret = 2
	if removeDuplicates(nums1) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ret = 5
	if removeDuplicates(nums2) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}

package leetcode

import (
	"testing"
)

func TestMerge(t *testing.T) {
	m, n := 3, 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	ret := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)
	for k, v := range ret {
		if nums1[k] != v {
			t.Fatalf("case fails: %v\n", nums1)
		}
	}
}

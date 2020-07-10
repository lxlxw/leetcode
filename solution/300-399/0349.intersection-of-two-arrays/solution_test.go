package leetcode

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	var nums1 []int
	var nums2 []int
	var ret []int

	nums1 = []int{1, 2, 2, 1}
	nums2 = []int{2, 2}
	ret = []int{2}
	for i, v := range intersection(nums1, nums2) {
		if v != ret[i] {
			t.Fatalf("case fails %v\n", ret)
		}
	}

	nums1 = []int{}
	nums2 = []int{2, 2}
	ret = []int{}
	fmt.Println(intersection(nums1, nums2))
	for i, v := range intersection(nums1, nums2) {
		if v != ret[i] {
			t.Fatalf("case fails %v\n", ret)
		}
	}
}

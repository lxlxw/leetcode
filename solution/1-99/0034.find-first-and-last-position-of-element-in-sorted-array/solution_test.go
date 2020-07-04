package leetcode

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	var nums []int
	var target int
	var ret []int

	target = 6
	nums = []int{5, 7, 7, 8, 8, 10}
	ret = []int{-1, -1}
	for i, v := range searchRange(nums, target) {
		if ret[i] != v {
			t.Fatalf("case fails %v\n", ret)
		}
	}

	target = 8
	nums = []int{5, 7, 7, 8, 8, 10}
	ret = []int{3, 4}
	for i, v := range searchRange(nums, target) {
		if ret[i] != v {
			t.Fatalf("case fails %v\n", ret)
		}
	}

}

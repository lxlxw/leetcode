package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	var target int
	var nums []int
	var ret []int

	nums = []int{2, 7, 11, 15}
	target = 9
	ret = []int{1, 2}

	for k, v := range twoSum(nums, target) {
		if v != ret[k] {
			t.Fatalf("case fails: %v\n", ret)
		}
	}

	nums = []int{2, 7, 11, 15}
	target = 13
	ret = []int{1, 3}

	for k, v := range twoSum(nums, target) {
		if v != ret[k] {
			t.Fatalf("case fails: %v\n", ret)
		}
	}
}

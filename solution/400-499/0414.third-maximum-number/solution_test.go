package leetcode

import (
	"testing"
)

func TestThirdMax(t *testing.T) {
	var ret int
	var nums []int

	nums = []int{3, 2, 1}
	ret = 1
	if ret != thirdMax(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	nums = []int{1, 2}
	ret = 2
	if ret != thirdMax(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	nums = []int{2, 1}
	ret = 2
	if ret != thirdMax(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	nums = []int{2, 2, 3, 1}
	ret = 1
	if ret != thirdMax(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	nums = []int{1}
	ret = 1
	if ret != thirdMax(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	nums = []int{2, 2, 3, 1, 5, 3, 5}
	ret = 2
	if ret != thirdMax(nums) {
		t.Fatalf("case fails %v\n", ret)
	}

}

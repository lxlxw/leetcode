package leetcode

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var s int
	var nums []int
	var ret int
	nums = []int{2, 3, 1, 2, 4, 3}
	ret = 2
	s = 7
	if ret != minSubArrayLen(s, nums) {
		t.Fatalf("case fails %v\n", ret)
	}

	nums = []int{}
	ret = 0
	s = 7
	if ret != minSubArrayLen(s, nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

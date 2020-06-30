package leetcode

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var ret int
	var nums []int
	ret = 7
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	if ret != removeDuplicates(nums) {
		t.Fatalf("case fails %v\n", ret)
	}
}

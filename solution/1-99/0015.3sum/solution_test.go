package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	var ret [][]int
	var nums []int

	nums = []int{-1, 0, 1, 2, -1, -4}
	ret = [][]int{{-1, -1, 2}, {-1, 0, 1}}
	for k, num := range threeSum(nums) {
		for i, v := range num {
			if v != ret[k][i] {
				t.Fatalf("case fails %v\n", ret)
			}
		}
	}
}

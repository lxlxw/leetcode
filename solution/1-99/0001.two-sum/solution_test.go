package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{1, 2, 3}
	target := 3
	result := []int{0, 1}

	fmt.Printf("nums = %v target = %v result = %v\n", nums, target, twoSum(nums, target))
	if ret := twoSum(nums, target); ret[0] != result[0] && ret[1] != result[1] {
		t.Fatalf("case fails: %v\n", ret)
	}

	nums = []int{1, 10, 20, 30}
	target = 30
	result = []int{1, 2}

	fmt.Printf("nums = %v target = %v result = %v\n", nums, target, twoSum(nums, target))
	if ret := twoSum(nums, target); ret[0] != result[0] && ret[1] != result[1] {
		t.Fatalf("case fails: %v\n", ret)
	}
}

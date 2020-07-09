package leetcode

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	var nums [][]int

	nums = [][]int{{1, 4}, {4, 5}}

	fmt.Println(merge(nums))

	nums = [][]int{{1, 3}}

	fmt.Println(merge(nums))

	nums = [][]int{{1, 4}, {5, 6}}

	fmt.Println(merge(nums))

	nums = [][]int{{1, 6}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println(merge(nums))

	nums = [][]int{{1, 4}, {0, 4}}

	fmt.Println(merge(nums))

	nums = [][]int{{1, 4}, {0, 1}}

	fmt.Println(merge(nums))

	nums = [][]int{{1, 4}, {0, 0}}

	fmt.Println(merge(nums))

	nums = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}

	fmt.Println(merge(nums))
}

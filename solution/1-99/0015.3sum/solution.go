package leetcode

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	flagMap := make(map[string]bool)
	var res [][]int
	i := 0
	for i < len(nums) {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		p, q := 1, len(nums)-1
		for p+i < q {
			if nums[i]+nums[p+i]+nums[q] == 0 {
				flag := strconv.Itoa(nums[i]) + strconv.Itoa(nums[p+i]) + strconv.Itoa(nums[q])
				if _, ok := flagMap[flag]; !ok {
					res = append(res, []int{nums[i], nums[p+i], nums[q]})
				}
				flagMap[flag] = true
				p++
				q--
			} else if nums[i]+nums[p+i]+nums[q] > 0 {
				q--
			} else if nums[i]+nums[p+i]+nums[q] < 9 {
				p++
			}
		}
		i++
	}
	return res
}

// @lc code=end

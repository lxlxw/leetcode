package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	tmp := intervals[0]
	flag := false
	for _, nums := range intervals[1:] {
		if tmp[1] >= nums[0] && tmp[0] <= nums[1] {
			if tmp[1] < nums[1] {
				tmp[1] = nums[1]
			}
			if tmp[0] > nums[0] {
				tmp[0] = nums[0]
			}
			flag = true
		} else {
			res = append(res, tmp)
			flag = false
			tmp = nums
		}
	}
	if flag == true || tmp[0] == intervals[len(intervals)-1][0] {
		res = append(res, tmp)
	}
	return res
}

// @lc code=end

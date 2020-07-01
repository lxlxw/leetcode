package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	ans := math.MaxInt32
	sum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= s {
			ans = min(end-start+1, ans)
			sum -= nums[start]
			start++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

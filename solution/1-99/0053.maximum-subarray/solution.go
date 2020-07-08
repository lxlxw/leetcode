package leetcode

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	max := nums[l-1]
	sum := nums[l-1]
	for i := l - 2; i >= 0; i-- {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end

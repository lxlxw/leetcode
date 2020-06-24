package leetcode

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	sum1, sum2 := len(nums), 0
	for i := range nums {
		sum1 += i
		sum2 += nums[i]
	}
	return sum1 - sum2
}

// @lc code=end

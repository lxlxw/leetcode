package leetcode

/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	p, q := 0, 0
	for i := range nums {
		if nums[i] == 1 {
			p++
			if q < p {
				q = p
			}
		} else {
			p = 0
		}
	}
	return q
}

// @lc code=end

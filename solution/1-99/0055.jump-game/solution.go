package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	max := 0
	for i := 0; i <= max; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return false
}

// @lc code=end

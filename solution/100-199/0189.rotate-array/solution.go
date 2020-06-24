package leetcode

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	length := len(nums)
	if length != 0 {
		i := length - k%length
		copy(nums, append(nums[i:length], nums[0:i]...))
	}
}

// @lc code=end

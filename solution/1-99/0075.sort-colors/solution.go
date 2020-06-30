package leetcode

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	for i := 0; i <= r; i++ {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
		if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			i--
			r--
		}
	}
}

// @lc code=end

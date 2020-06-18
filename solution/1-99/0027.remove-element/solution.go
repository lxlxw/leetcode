package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	lenNum := len(nums)
	if lenNum == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...) // 移除该元素
			i--
			lenNum--
		}
	}
	return lenNum
}

// @lc code=end

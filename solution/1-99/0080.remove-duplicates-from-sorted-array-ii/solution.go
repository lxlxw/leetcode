package leetcode

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	numMap := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		if numMap[nums[i]] >= 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			l--
		}
		numMap[nums[i]]++
	}
	return len(nums)
}

// @lc code=end

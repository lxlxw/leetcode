package leetcode

/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	var numMap = make(map[int]int, 0)
	for k, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = k
	}
	return false
}

// @lc code=end

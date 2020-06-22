package leetcode

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	var numMap = make(map[int]int, 0)
	for i, v := range nums {
		if ii, ok := numMap[v]; ok && (i-ii) <= k {
			return true
		}
		numMap[v] = i
	}
	return false
}

// @lc code=end

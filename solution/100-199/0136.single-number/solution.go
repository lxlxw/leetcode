package leetcode

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	numMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			numMap[v] = false
		} else {
			numMap[v] = true
		}
	}
	for _, v := range nums {
		if numMap[v] == true {
			return v
		}
	}
	return 0
}

// @lc code=end

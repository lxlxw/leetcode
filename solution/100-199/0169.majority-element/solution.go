package leetcode

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	n := l/2 + l%2 // l%2为额外加个奇偶数的判断
	numMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			numMap[v]++
		} else {
			numMap[v] = 1
		}
		if numMap[v] >= n {
			return v
		}
	}
	return -1
}

// @lc code=end

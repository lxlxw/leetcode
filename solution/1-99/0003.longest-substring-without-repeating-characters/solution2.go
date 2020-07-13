package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range s {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(s)-left > max {
		max = len(s) - left
	}
	return max
}

// @lc code=end

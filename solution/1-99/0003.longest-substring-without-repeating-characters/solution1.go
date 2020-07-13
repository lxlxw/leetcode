package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring1(s string) int {
	var res int
	for i := range s {
		sMap := make(map[byte]bool)
		q := 0
		for q+i < len(s) {
			if sMap[s[q+i]] {
				break
			} else {
				sMap[s[q+i]] = true
			}
			q++
		}
		if q > res {
			res = q
		}
		if q >= len(s)-i {
			break
		}
	}
	return res
}

// @lc code=end

package leetcode

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	p := 0
	q := 0
	for q < len(needle) && p < len(haystack) {
		if haystack[p:p+1] == needle[q:q+1] {
			q++
		} else {
			p = p - q
			q = 0
		}
		p++
	}
	if q != len(needle) {
		return -1
	}
	return p - q
}

// @lc code=end

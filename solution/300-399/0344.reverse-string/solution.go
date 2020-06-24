package leetcode

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	p, q := 0, len(s)-1
	for p < q {
		s[p], s[q] = s[q], s[p]
		p++
		q--
	}
}

// @lc code=end

package leetcode

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	b := []byte(s)
	l := 0
	for i, v := range s {
		if v == ' ' || i == len(s)-1 {
			r := i - 1
			if i == len(s)-1 {
				r = i
			}
			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}
			l = i + 1
		}
	}
	return string(b)
}

// @lc code=end

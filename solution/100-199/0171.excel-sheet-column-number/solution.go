package leetcode

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	res := 0
	for _, char := range s {
		res *= 26
		res += int(char-'A') + 1
	}
	return res
}

// @lc code=end

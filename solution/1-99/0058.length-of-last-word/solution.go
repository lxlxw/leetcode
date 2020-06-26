package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")
	return len(sArr[len(sArr)-1])
}

// @lc code=end

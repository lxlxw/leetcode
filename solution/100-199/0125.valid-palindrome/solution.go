package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9') {
			l++
		}
		for l < r && (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9') {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

// @lc code=end

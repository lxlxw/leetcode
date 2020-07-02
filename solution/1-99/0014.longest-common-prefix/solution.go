package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res string
	sort.Strings(strs)
	for k := range strs[0] {
		flag := true
		for i := 0; i < len(strs)-1; i++ {
			if strs[i][k] != strs[i+1][k] {
				flag = false
				break
			}
		}
		if flag {
			res += strs[0][k : k+1]
		} else {
			break
		}
	}
	return res
}

// @lc code=end

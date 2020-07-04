package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {

	sArr := strings.Split(str, " ")
	if len(sArr) != len(pattern) {
		return false
	}
	pMap := make(map[byte]string)
	qMap := make(map[string]byte)

	for i, v := range sArr {
		if _, ok := pMap[pattern[i]]; ok && pMap[pattern[i]] != v {
			return false
		} else {
			pMap[pattern[i]] = v
		}
		if _, ok := qMap[v]; ok && qMap[v] != pattern[i] {
			return false
		} else {
			qMap[v] = pattern[i]
		}
	}
	return true
}

// @lc code=end

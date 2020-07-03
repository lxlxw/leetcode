package leetcode

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	sMap, tMap := map[uint8]int{}, map[uint8]int{}
	for i := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		} else {
			sMap[s[i]] = i + 1
			tMap[t[i]] = i + 1
		}
	}
	return true
}

// @lc code=end

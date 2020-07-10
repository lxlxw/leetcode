package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	for i := range s {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]]++
		} else {
			sMap[s[i]] = 1
		}
	}

	for i := range t {
		if _, ok := sMap[t[i]]; ok && sMap[t[i]] > 0 {
			sMap[t[i]]--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

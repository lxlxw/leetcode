package leetcode

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	sMap := make(map[string]bool)
	var res []string
	for i := 0; i < len(s)-9; i++ {
		if _, ok := sMap[s[i:i+10]]; ok {
			if sMap[s[i:i+10]] == true {
				res = append(res, s[i:i+10])
			}
			sMap[s[i:i+10]] = false
		} else {
			sMap[s[i:i+10]] = true
		}
	}
	return res
}

// @lc code=end

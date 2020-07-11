/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {

	// abccccdd
	sMap := make(map[byte]int)
	for i := range s {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]]++
		} else {
			sMap[s[i]] = 1
		}
	}
	var count int
	var flag bool
	for _, num := range sMap {
		if num%2 != 0 {
			flag = true
			num--
		}
		count += num
	}
	if flag == true {
		count++
	}
	return count
}
// @lc code=end


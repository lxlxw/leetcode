package leetcode

import "unicode"

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	vowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	p, q := 0, len(s)-1
	ret := []byte(s)
	for p < q {
		_, ok1 := vowel[unicode.ToLower(rune(ret[p]))]
		if !ok1 {
			p++
		}
		_, ok2 := vowel[unicode.ToLower(rune(ret[q]))]
		if !ok2 {
			q--
		}
		if ok1 && ok2 {
			ret[p], ret[q] = ret[q], ret[p]
			p++
			q--
		}
	}
	return string(ret)
}

// @lc code=end

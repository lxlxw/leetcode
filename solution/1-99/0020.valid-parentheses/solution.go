package leetcode

import (
	"container/list"
	"strings"
)

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if s == "" {
		return true
	}
	s = strings.Replace(s, " ", "", -1)
	stack := list.New()
	trMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for v := range s {
		switch s[v : v+1] {
		case "(", "{", "[":
			stack.PushBack(s[v : v+1])
		case ")", "}", "]":
			if stack.Len() == 0 || stack.Back().Value != trMap[s[v:v+1]] {
				return false
			} else {
				stack.Remove(stack.Back())
			}
		}
	}
	return stack.Len() == 0
}

// @lc code=end

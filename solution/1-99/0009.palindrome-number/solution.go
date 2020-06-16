package leetcode

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	len := len(str)
	for i := 0; i < len; i++ {
		if str[i:i+1] != str[len-i-1:len-i] {
			return false
		}
	}
	return true
}

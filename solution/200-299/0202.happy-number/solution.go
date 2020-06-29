package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	numMap := make(map[int]bool)
	for n != 1 {
		num := 0
		for i := 0; i < len(strconv.Itoa(n)); i++ {
			v := n / powerf(10, i) % 10
			num += v * v
		}
		n = num
		if _, ok := numMap[n]; ok {
			return false
		}
		numMap[n] = true
	}
	return true
}

func powerf(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * powerf(x, n-1)
}

// @lc code=end

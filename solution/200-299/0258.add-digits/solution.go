package leetcode

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num != 0 {
		sum += num % 10
		num = num / 10
	}
	return addDigits(sum)
}

// @lc code=end

package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x
	for {
		mid := (l + r) / 2
		if mid == l {
			return mid
		}
		sqr := mid * mid
		if sqr > x {
			r = mid
		}
		if sqr < x {
			l = mid
		}
		if sqr == x {
			return mid
		}
	}
}

// @lc code=end

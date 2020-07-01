package leetcode

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00000
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return myPow(1.0/x, -n)
	}
	if n%2 != 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

// @lc code=end

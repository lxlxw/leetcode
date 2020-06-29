package leetcode

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	var count int
	fMap := make(map[int]bool)
	for i := 2; i < n; i++ {
		if fMap[i] {
			continue
		}
		count++
		fMap[i] = false
		// i的倍数不可能为素数
		for j := 2 * i; j < n; j += i {
			fMap[j] = true
		}
	}
	return count
}

// @lc code=end

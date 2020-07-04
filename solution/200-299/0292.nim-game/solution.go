package leetcode

/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
func canWinNim(n int) bool {
	if n%4 != 0 {
		return true
	}
	return false
}

// @lc code=end

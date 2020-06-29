package leetcode

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	// ascii 1-9 = 48-57
	var row, col, block [9][58]bool
	for k, nums := range board {
		for i, v := range nums {
			if v == '.' {
				continue
			}
			if row[i][v] {
				return false
			}
			if col[k][v] {
				return false
			}
			if block[i/3+k/3*3][v] {
				return false
			}
			row[i][v] = true
			col[k][v] = true
			block[i/3+k/3*3][v] = true
		}
	}
	return true
}

// @lc code=end

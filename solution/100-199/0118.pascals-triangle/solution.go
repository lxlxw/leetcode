package leetcode

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	var res [][]int
	res = append(res, []int{1})
	res = append(res, []int{1, 1})
	recursive(2, numRows, &res)
	return res
}

func recursive(current, numRows int, res *[][]int) {
	row := make([]int, current+1)
	row[0], row[current] = 1, 1
	for i := 1; i < current; i++ {
		row[i] = (*res)[current-1][i-1] + (*res)[current-1][i]
	}
	*res = append(*res, row)
	if current+1 == numRows {
		return
	} else {
		recursive(current+1, numRows, res)
	}
}

// @lc code=end

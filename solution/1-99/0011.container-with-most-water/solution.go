package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	p, q := 0, len(height)-1
	var ret int
	var tmp int
	for p < q {
		if height[p] < height[q] {
			tmp = height[p] * (q - p)
			p++
		} else {
			tmp = height[q] * (q - p)
			q--
		}
		if ret < tmp {
			ret = tmp
		}
	}
	return ret
}

// @lc code=end

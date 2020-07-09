package leetcode

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	sale, max := 0, 0
	for _, v := range prices[1:] {
		if v-buy > 0 {
			sale = v - buy
		}
		if sale > max {
			max = sale
		}
		if v < buy {
			buy = v
		}
	}
	return max
}

// @lc code=end

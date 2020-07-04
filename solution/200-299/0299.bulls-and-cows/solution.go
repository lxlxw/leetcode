package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	var countS, countG [10]int
	bulls, cows := 0, 0
	for i := range secret {
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')
		if ng == ns {
			bulls++
			continue
		}

		if countG[ns] > 0 {
			cows++
			countG[ns]--
		} else {
			countS[ns]++
		}

		if countS[ng] > 0 {
			cows++
			countS[ng]--
		} else {
			countG[ng]++
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

// @lc code=end

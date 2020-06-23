package leetcode

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left != right {
		if numbers[left]+numbers[right] == target {
			break
		}
		if numbers[left]+numbers[right] > target {
			right--
		}
		if numbers[left]+numbers[right] < target {
			left++
		}
	}
	return []int{left + 1, right + 1}
}

// @lc code=end

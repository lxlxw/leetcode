package leetcode

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			res := []int{mid, mid}
			for res[0]-1 >= 0 && nums[res[0]-1] == target {
				res[0]--
			}
			for res[1]+1 < len(nums) && nums[res[1]+1] == target {
				res[1]++
			}
			return res
		}
	}

	return []int{-1, -1}
}

// @lc code=end

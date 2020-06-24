package leetcode

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	first, second, third := ^1<<32, ^1<<32, ^1<<32
	for i := range nums {
		if first == nums[i] || second == nums[i] || third == nums[i] {
			continue
		}
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}
	}
	if third == ^1<<32 {
		return first
	}
	return third
}

// @lc code=end

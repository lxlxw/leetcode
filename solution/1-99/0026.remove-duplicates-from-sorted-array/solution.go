package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p := 1
	for q := 1; q < len(nums); q++ {
		if nums[q-1] != nums[q] {
			nums[p] = nums[q]
			p++
		}
	}
	return p
}

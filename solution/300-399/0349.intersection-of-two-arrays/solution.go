package leetcode

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	num1Map := make(map[int]bool)

	for i := range nums1 {
		num1Map[nums1[i]] = true
	}
	num2Map := make(map[int]bool)
	var res []int
	for i := range nums2 {
		if _, ok1 := num1Map[nums2[i]]; ok1 && !num2Map[nums2[i]] {
			num2Map[nums2[i]] = true
			res = append(res, nums2[i])
		}
	}
	return res
}

// @lc code=end

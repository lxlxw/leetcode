package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j = m - 1, n - 1

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

// @lc code=end

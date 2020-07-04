# [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array)

### 题目描述

<p>给定一个按照升序排列的整数数组 <code>nums</code>，和一个目标值 <code>target</code>。找出给定目标值在数组中的开始位置和结束位置。</p>

<p>你的算法时间复杂度必须是&nbsp;<em>O</em>(log <em>n</em>) 级别。</p>

<p>如果数组中不存在目标值，返回&nbsp;<code>[-1, -1]</code>。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> nums = [<code>5,7,7,8,8,10]</code>, target = 8
<strong>输出:</strong> [3,4]</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> nums = [<code>5,7,7,8,8,10]</code>, target = 6
<strong>输出:</strong> [-1,-1]</pre>



### 解题思路


### 具体解法

#### **Golang**
```go
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
```



# [349. 两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays/description/)

### 题目描述

<p>给定两个数组，编写一个函数来计算它们的交集。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums1 = [1,2,2,1], nums2 = [2,2]
<strong>输出：</strong>[2]
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums1 = [4,9,5], nums2 = [9,4,9,8,4]
<strong>输出：</strong>[9,4]</pre>

<p>&nbsp;</p>

<p><strong>说明：</strong></p>

<ul>
	<li>输出结果中的每个元素一定是唯一的。</li>
	<li>我们可以不考虑输出结果的顺序。</li>
</ul>

### 解题思路

1. hash table

### 具体解法


#### **Golang**
```go
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
```



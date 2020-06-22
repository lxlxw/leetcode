# [219.存在重复元素 II](https://leetcode-cn.com/problems/contains-duplicate-ii/description/)


### 题目描述

<div class="notranslate"><p>给定一个整数数组和一个整数&nbsp;<em>k</em>，判断数组中是否存在两个不同的索引<em>&nbsp;i</em>&nbsp;和<em>&nbsp;j</em>，使得&nbsp;<strong>nums [i] = nums [j]</strong>，并且 <em>i</em> 和 <em>j</em>&nbsp;的差的 <strong>绝对值</strong> 至多为 <em>k</em>。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> nums = [1,2,3,1], k<em> </em>= 3
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>nums = [1,0,1,1], k<em> </em>=<em> </em>1
<strong>输出:</strong> true</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入: </strong>nums = [1,2,3,1,2,3], k<em> </em>=<em> </em>2
<strong>输出:</strong> false</pre>
</div>

### 解题思路

1. 利用hash表来实现

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func containsNearbyDuplicate(nums []int, k int) bool {
	var numMap = make(map[int]int, 0)
	for i, v := range nums {
		if ii, ok := numMap[v]; ok && (i-ii) <= k {
			return true
		}
		numMap[v] = i
	}
	return false
}
```


<!-- tabs:end -->
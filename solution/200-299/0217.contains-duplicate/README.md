# [217.存在重复元素](https://leetcode-cn.com/problems/contains-duplicate/)


### 题目描述

<div class="notranslate"><p>给定一个整数数组，判断是否存在重复元素。</p>

<p>如果任意一值在数组中出现至少两次，函数返回 <code>true</code> 。如果数组中每个元素都不相同，则返回 <code>false</code> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [1,2,3,1]
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>[1,2,3,4]
<strong>输出:</strong> false</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入: </strong>[1,1,1,3,3,4,3,2,4,2]
<strong>输出:</strong> true</pre>
</div>

### 解题思路

1. 利用hash表来实现

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func containsDuplicate(nums []int) bool {
	var numMap = make(map[int]int, 0)
	for k, v := range nums {
		if _, ok := numMap[v]; ok {
			return true
		}
		numMap[v] = k
	}
	return false
}
```


<!-- tabs:end -->
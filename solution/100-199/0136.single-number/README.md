# [136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/description/)

### 题目描述

<p>给定一个<strong>非空</strong>整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。</p>

<p><strong>说明：</strong></p>

<p>你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [2,2,1]
<strong>输出:</strong> 1
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [4,1,2,1,2]
<strong>输出:</strong> 4</pre>

### 解题思路

1. hash table

### 具体解法

#### **Golang**
```go
func singleNumber(nums []int) int {
	numMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			numMap[v] = false
		} else {
			numMap[v] = true
		}
	}
	for _, v := range nums {
		if numMap[v] == true {
			return v
		}
	}
	return 0
}
```



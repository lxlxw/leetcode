# [169.多数元素](https://leetcode-cn.com/problems/majority-element/description/)


### 题目描述

<div class="notranslate"><p>给定一个大小为 <em>n </em>的数组，找到其中的多数元素。多数元素是指在数组中出现次数<strong>大于</strong>&nbsp;<code>⌊ n/2 ⌋</code>&nbsp;的元素。</p>

<p>你可以假设数组是非空的，并且给定的数组总是存在多数元素。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> [3,2,3]
<strong>输出:</strong> 3</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [2,2,1,1,1,2,2]
<strong>输出:</strong> 2
</pre>
</div>

### 解题思路

1. 利用hash表来实现

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func majorityElement(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	n := l/2 + l%2 // l%2为额外加个奇偶数的判断
	numMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			numMap[v]++
		} else {
			numMap[v] = 1
		}
		if numMap[v] >= n {
			return v
		}
	}
	return -1
}
```


<!-- tabs:end -->
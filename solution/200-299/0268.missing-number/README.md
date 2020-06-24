# [268.缺失数字](https://leetcode-cn.com/problems/missing-number/description/)


### 题目描述

<div class="notranslate"><p>给定一个包含 <code>0, 1, 2, ..., n</code>&nbsp;中&nbsp;<em>n</em>&nbsp;个数的序列，找出 0 .. <em>n</em>&nbsp;中没有出现在序列中的那个数。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [3,0,1]
<strong>输出:</strong> 2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [9,6,4,2,3,5,7,0,1]
<strong>输出:</strong> 8
</pre>

<p>&nbsp;</p>

<p><strong>说明:</strong><br>
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?</p>
</div>


### 解题思路

1. `sum1`为 `0, 1, 2, ..., n`所有的和
2. `sum2`为数组中所有值得和
3. 两者相减即可得到缺失的数

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func missingNumber(nums []int) int {
	sum1, sum2 := len(nums), 0
	for i, _ := range nums {
		sum1 += i
		sum2 += nums[i]
	}
	return sum1 - sum2
}
```


<!-- tabs:end -->
# [414. 第三大的数](https://leetcode-cn.com/problems/third-maximum-number/description/)


### 题目描述

<p>给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> [3, 2, 1]

<strong>输出:</strong> 1

<strong>解释:</strong> 第三大的数是 1.
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> [1, 2]

<strong>输出:</strong> 2

<strong>解释:</strong> 第三大的数不存在, 所以返回最大的数 2 .
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> [2, 2, 3, 1]

<strong>输出:</strong> 1

<strong>解释:</strong> 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。
</pre>

### 解题思路

1. 用三个变量来存放第一大，第二大，第三大的元素的变量，分别为first, second, third
2. 遍历数组，若该元素比first大则往后顺移一个元素，比second大则往往后顺移一个元素，比third大则赋值个third
3. 最后得到第三大的元素，若没有则返回第一大的元素。


### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	first, second, third := ^1<<32, ^1<<32, ^1<<32
	for i := range nums {
		if first == nums[i] || second == nums[i] || third == nums[i] {
			continue
		}
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}
	}
	if third == ^1<<32 {
		return first
	}
	return third
}
```


<!-- tabs:end -->
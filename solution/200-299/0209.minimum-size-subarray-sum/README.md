# [209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/)

### 题目描述

<p>给定一个含有&nbsp;<strong>n&nbsp;</strong>个正整数的数组和一个正整数&nbsp;<strong>s ，</strong>找出该数组中满足其和<strong> &ge; s </strong>的长度最小的连续子数组，并返回其长度<strong>。</strong>如果不存在符合条件的连续子数组，返回 0。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong><code>s = 7, nums = [2,3,1,2,4,3]</code>
<strong>输出：</strong>2
<strong>解释：</strong>子数组&nbsp;<code>[4,3]</code>&nbsp;是该条件下的长度最小的连续子数组。
</pre>

<p>&nbsp;</p>

<p><strong>进阶：</strong></p>

<ul>
	<li>如果你已经完成了<em> O</em>(<em>n</em>) 时间复杂度的解法, 请尝试 <em>O</em>(<em>n</em> log <em>n</em>) 时间复杂度的解法。</li>
</ul>

### 解题思路

定义两个指针 start 和 end 分别表示子数组的开始位置和结束位置，维护变量 sum 存储子数组中的元素和（即从 []nums[start] 到 []nums[end] 的元素和）。

初始状态下，start 和 end 都指向下标 00，sum 的值为 00。

每一轮迭代，将 [end]nums[end] 加到 sum，如果  \ge ssum≥s，则更新子数组的最小长度（此时子数组的长度是 -+1end−start+1），然后将 [start]nums[start] 从 sum 中减去并将 start 右移，直到 sum < s，在此过程中同样更新子数组的最小长度。在每一轮迭代的最后，将 end 右移。


### 具体解法


#### **Golang**
```go
func minSubArrayLen(s int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	ans := math.MaxInt32
	sum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= s {
			ans = min(end-start+1, ans)
			sum -= nums[start]
			start++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
```



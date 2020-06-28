# [15. 三数之和](https://leetcode-cn.com/problems/3sum/description/)

### 题目描述

<p>给你一个包含 <em>n</em> 个整数的数组&nbsp;<code>nums</code>，判断&nbsp;<code>nums</code>&nbsp;中是否存在三个元素 <em>a，b，c ，</em>使得&nbsp;<em>a + b + c = </em>0 ？请你找出所有满足条件且不重复的三元组。</p>

<p><strong>注意：</strong>答案中不可以包含重复的三元组。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre>给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
</pre>

### 解题思路

![](http://lc-photo.xwlin.com/15.gif)

### 具体解法

#### **Golang**
```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	flagMap := make(map[string]bool)
	var res [][]int
	i := 0
	for i < len(nums) {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		p, q := 1, len(nums)-1
		for p+i < q {
			if nums[i]+nums[p+i]+nums[q] == 0 {
				flag := strconv.Itoa(nums[i]) + strconv.Itoa(nums[p+i]) + strconv.Itoa(nums[q])
				if _, ok := flagMap[flag]; !ok {
					res = append(res, []int{nums[i], nums[p+i], nums[q]})
				}
				flagMap[flag] = true
				p++
				q--
			} else if nums[i]+nums[p+i]+nums[q] > 0 {
				q--
			} else if nums[i]+nums[p+i]+nums[q] < 9 {
				p++
			}
		}
		i++
	}
	return res
}
```


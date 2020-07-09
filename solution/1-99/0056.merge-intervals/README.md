# [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals)

### 题目描述

<p>给出一个区间的集合，请合并所有重叠的区间。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [[1,3],[2,6],[8,10],[15,18]]
<strong>输出:</strong> [[1,6],[8,10],[15,18]]
<strong>解释:</strong> 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [[1,4],[4,5]]
<strong>输出:</strong> [[1,5]]
<strong>解释:</strong> 区间 [1,4] 和 [4,5] 可被视为重叠区间。</pre>



### 解题思路


### 具体解法


#### **Golang**
```go
func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	tmp := intervals[0]
	flag := false
	for _, nums := range intervals[1:] {
		if tmp[1] >= nums[0] && tmp[0] <= nums[1] {
			if tmp[1] < nums[1] {
				tmp[1] = nums[1]
			}
			if tmp[0] > nums[0] {
				tmp[0] = nums[0]
			}
			flag = true
		} else {
			res = append(res, tmp)
			flag = false
			tmp = nums
		}
	}
	if flag == true || tmp[0] == intervals[len(intervals)-1][0] {
		res = append(res, tmp)
	}
	return res
}
```



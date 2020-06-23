# [283.移动零](https://leetcode-cn.com/problems/move-zeroes/description/)


### 题目描述

<div class="notranslate"><p>给定一个数组 <code>nums</code>，编写一个函数将所有 <code>0</code> 移动到数组的末尾，同时保持非零元素的相对顺序。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>[0,1,0,3,12]</code>
<strong>输出:</strong> <code>[1,3,12,0,0]</code></pre>

<p><strong>说明</strong>:</p>

<ol>
	<li>必须在原数组上操作，不能拷贝额外的数组。</li>
	<li>尽量减少操作次数。</li>
</ol>
</div>


### 解题思路

![](http://lc-photo.xwlin.com/283.gif)

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func moveZeroes(nums []int) {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p] = nums[i]
			if p != i {
				nums[i] = 0
			}
			p++
		}
	}
}
```


<!-- tabs:end -->
# [88.合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array)


### 题目描述

<div class="notranslate"><p>给你两个有序整数数组&nbsp;<em>nums1 </em>和 <em>nums2</em>，请你将 <em>nums2 </em>合并到&nbsp;<em>nums1&nbsp;</em>中<em>，</em>使 <em>nums1 </em>成为一个有序数组。</p>

<p>&nbsp;</p>

<p><strong>说明:</strong></p>

<ul>
	<li>初始化&nbsp;<em>nums1</em> 和 <em>nums2</em> 的元素数量分别为&nbsp;<em>m</em> 和 <em>n </em>。</li>
	<li>你可以假设&nbsp;<em>nums1&nbsp;</em>有足够的空间（空间大小大于或等于&nbsp;<em>m + n</em>）来保存 <em>nums2</em> 中的元素。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

<strong>输出:</strong>&nbsp;[1,2,2,3,5,6]</pre>
</div>

### 解题思路

1. 采用双指针分别将两个数组 `nums1` 和数组 `nums2` 从后向前插入到数组 `nums1` 中
2. 指针 `i` 标记数组 nums1 当前比较位置，开始位数组元素最后位
3. 指针 `j` 标记数组 nums2 当前比较位置，开始位数组元素最后位
4. 指针 `k` 表示结果从后向前当前可插入位置，默认 `nums1` 后 `n` 位都是可插入位，将 `nums1` 中元素插入当前可插入位后， 从 指针 `i` 至 `k` 都是可插入位
5. 循环将数组 `nums2` 插入到数组 `nums1`
6. 当数组 `nums1` 当前可比较数据位 `i` 大于等于 0 并且数组 `nums1` 当前位置值比 `nums2` 当前位置值大时，将数组 `nums1` 当前位置数组 插入到 `k` 位置 ，`i` 位指针向前挪一位
7. 否则将 `nums2` 当前位置数据插入到 `k` 位置，`j` 位指针向前挪一位
8. 插入完成后将当前插入位置指针 `k` 向前挪一位，进行下一次插入

![](http://lc-photo.xwlin.com/88.png)

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j = m - 1, n - 1

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
```


<!-- tabs:end -->
# [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx)

### 题目描述

<p>实现&nbsp;<code>int sqrt(int x)</code>&nbsp;函数。</p>

<p>计算并返回&nbsp;<em>x</em>&nbsp;的平方根，其中&nbsp;<em>x </em>是非负整数。</p>

<p>由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 4
<strong>输出:</strong> 2
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> 8
<strong>输出:</strong> 2
<strong>说明:</strong> 8 的平方根是 2.82842..., 
&nbsp;    由于返回类型是整数，小数部分将被舍去。
</pre>



### 解题思路

1. 二分查找

### 具体解法


#### **Golang**
```go
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x
	for {
		mid := (l + r) / 2
		if mid == l {
			return mid
		}
		sqr := mid * mid
		if sqr > x {
			r = mid
		}
		if sqr < x {
			l = mid
		}
		if sqr == x {
			return mid
		}
	}
}

```



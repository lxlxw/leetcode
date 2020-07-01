# [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n)

### 题目描述

<p>实现&nbsp;<a href="https://www.cplusplus.com/reference/valarray/pow/" target="_blank">pow(<em>x</em>, <em>n</em>)</a>&nbsp;，即计算 x 的 n 次幂函数。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> 2.00000, 10
<strong>输出:</strong> 1024.00000
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 2.10000, 3
<strong>输出:</strong> 9.26100
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong> 2.00000, -2
<strong>输出:</strong> 0.25000
<strong>解释:</strong> 2<sup>-2</sup> = 1/2<sup>2</sup> = 1/4 = 0.25</pre>

<p><strong>说明:</strong></p>

<ul>
	<li>-100.0 &lt;&nbsp;<em>x</em>&nbsp;&lt; 100.0</li>
	<li><em>n</em>&nbsp;是 32 位有符号整数，其数值范围是&nbsp;[&minus;2<sup>31</sup>,&nbsp;2<sup>31&nbsp;</sup>&minus; 1] 。</li>
</ul>


### 解题思路

1. 递归解法
2. 当n等于偶数时，特殊处理下

### 具体解法


#### **Golang**
```go
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00000
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return myPow(1.0/x, -n)
	}
	if n%2 != 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}
```

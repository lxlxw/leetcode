# [204. 计数质数](https://leetcode-cn.com/problems/count-primes/description/)

### 题目描述

<p>统计所有小于非负整数&nbsp;<em>n&nbsp;</em>的质数的数量。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 10
<strong>输出:</strong> 4
<strong>解释:</strong> 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
</pre>

### 解题思路

1. hash table
2. i的倍数不可能为素数

### 具体解法


#### **Golang**
```go
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	var count int
	fMap := make(map[int]bool)
	for i := 2; i < n; i++ {
		if fMap[i] {
			continue
		}
		count++
		fMap[i] = false
		// i的倍数不可能为素数
		for j := 2 * i; j < n; j += i {
			fMap[j] = true
		}
	}
	return count
}
```



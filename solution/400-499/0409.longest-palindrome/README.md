# [409. 最长回文串]https://leetcode-cn.com/problems/longest-palindrome/description/)

### 题目描述

<p>给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。</p>

<p>在构造过程中，请注意区分大小写。比如&nbsp;<code>&quot;Aa&quot;</code>&nbsp;不能当做一个回文字符串。</p>

<p><strong>注意:</strong><br />
假设字符串的长度不会超过 1010。</p>

<p><strong>示例 1: </strong></p>

<pre>
输入:
&quot;abccccdd&quot;

输出:
7

解释:
我们可以构造的最长的回文串是&quot;dccaccd&quot;, 它的长度是 7。
</pre>

### 解题思路


### 具体解法


#### **Golang**
```go
func longestPalindrome(s string) int {

	// abccccdd
	sMap := make(map[byte]int)
	for i := range s {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]]++
		} else {
			sMap[s[i]] = 1
		}
	}
	var count int
	var flag bool
	for _, num := range sMap {
		if num%2 != 0 {
			flag = true
			num--
		}
		count += num
	}
	if flag == true {
		count++
	}
	return count
}
```


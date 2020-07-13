# [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters)

### 题目描述

<p>给定一个字符串，请你找出其中不含有重复字符的&nbsp;<strong>最长子串&nbsp;</strong>的长度。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入: </strong>&quot;abcabcbb&quot;
<strong>输出: </strong>3 
<strong>解释:</strong> 因为无重复字符的最长子串是 <code>&quot;abc&quot;，所以其</code>长度为 3。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>&quot;bbbbb&quot;
<strong>输出: </strong>1
<strong>解释: </strong>因为无重复字符的最长子串是 <code>&quot;b&quot;</code>，所以其长度为 1。
</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入: </strong>&quot;pwwkew&quot;
<strong>输出: </strong>3
<strong>解释: </strong>因为无重复字符的最长子串是&nbsp;<code>&quot;wke&quot;</code>，所以其长度为 3。
&nbsp;    请注意，你的答案必须是 <strong>子串 </strong>的长度，<code>&quot;pwke&quot;</code>&nbsp;是一个<em>子序列，</em>不是子串。
</pre>



### 解题思路


### 具体解法

<!-- tabs:start -->

#### **Golang1**
```go
func lengthOfLongestSubstring1(s string) int {
	var res int
	for i := range s {
		sMap := make(map[byte]bool)
		q := 0
		for q+i < len(s) {
			if sMap[s[q+i]] {
				break
			} else {
				sMap[s[q+i]] = true
			}
			q++
		}
		if q > res {
			res = q
		}
		if q >= len(s)-i {
			break
		}
	}
	return res
}
```

#### **Golang2**
```go
func lengthOfLongestSubstring(s string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range s {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(s)-left > max {
		max = len(s) - left
	}
	return max
}
```

<!-- tabs:end -->


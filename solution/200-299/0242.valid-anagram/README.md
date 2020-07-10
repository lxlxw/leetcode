# [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/description/)

### 题目描述

<p>给定两个字符串 <em>s</em> 和 <em>t</em> ，编写一个函数来判断 <em>t</em> 是否是 <em>s</em> 的字母异位词。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> <em>s</em> = &quot;anagram&quot;, <em>t</em> = &quot;nagaram&quot;
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> <em>s</em> = &quot;rat&quot;, <em>t</em> = &quot;car&quot;
<strong>输出: </strong>false</pre>

<p><strong>说明:</strong><br>
你可以假设字符串只包含小写字母。</p>

<p><strong>进阶:</strong><br>
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？</p>

### 解题思路

1. hash table

### 具体解法


#### **Golang**
```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	for i := range s {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]]++
		} else {
			sMap[s[i]] = 1
		}
	}

	for i := range t {
		if _, ok := sMap[t[i]]; ok && sMap[t[i]] > 0 {
			sMap[t[i]]--
		} else {
			return false
		}
	}
	return true
}
```



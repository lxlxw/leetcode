# [205. 同构字符串](https://leetcode-cn.com/problems/isomorphic-strings/description/)

### 题目描述

<p>给定两个字符串&nbsp;<em><strong>s&nbsp;</strong></em>和&nbsp;<strong><em>t</em></strong>，判断它们是否是同构的。</p>

<p>如果&nbsp;<em><strong>s&nbsp;</strong></em>中的字符可以被替换得到&nbsp;<strong><em>t&nbsp;</em></strong>，那么这两个字符串是同构的。</p>

<p>所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> <strong><em>s</em></strong> = <code>&quot;egg&quot;, </code><strong><em>t = </em></strong><code>&quot;add&quot;</code>
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> <strong><em>s</em></strong> = <code>&quot;foo&quot;, </code><strong><em>t = </em></strong><code>&quot;bar&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> <strong><em>s</em></strong> = <code>&quot;paper&quot;, </code><strong><em>t = </em></strong><code>&quot;title&quot;</code>
<strong>输出:</strong> true</pre>

<p><strong>说明:</strong><br>
你可以假设&nbsp;<em><strong>s&nbsp;</strong></em>和 <strong><em>t </em></strong>具有相同的长度。</p>

### 解题思路


### 具体解法

1. 设置`s`和`t`的`hash table`
2. 遍历随意一个字符串
3. 将遍历到的字符串的当作`key`, 字符串下标 `i+1` 当作`value`保存在`map`中
4. 判断`map`中相同下标的字符串的对应下标位置是否相同

#### **Golang**
```go
func isIsomorphic(s string, t string) bool {
	sMap, tMap := map[uint8]int{}, map[uint8]int{}
	for i := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		} else {
			sMap[s[i]] = i + 1
			tMap[t[i]] = i + 1
		}
	}
	return true
}
```



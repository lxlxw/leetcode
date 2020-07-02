# [14. 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix)

### 题目描述

<p>编写一个函数来查找字符串数组中的最长公共前缀。</p>

<p>如果不存在公共前缀，返回空字符串&nbsp;<code>&quot;&quot;</code>。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入: </strong>[&quot;flower&quot;,&quot;flow&quot;,&quot;flight&quot;]
<strong>输出:</strong> &quot;fl&quot;
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入: </strong>[&quot;dog&quot;,&quot;racecar&quot;,&quot;car&quot;]
<strong>输出:</strong> &quot;&quot;
<strong>解释:</strong> 输入不存在公共前缀。
</pre>

<p><strong>说明:</strong></p>

<p>所有输入只包含小写字母&nbsp;<code>a-z</code>&nbsp;。</p>



### 解题思路

![](http://lc-photo.xwlin.com/14.png)

1. 纵向扫描。纵向扫描时，从前往后遍历所有字符串的每一列，比较相同列上的字符是否相同，如果相同则继续对下一列进行比较，如果不相同则当前列不再属于公共前缀，当前列之前的部分为最长公共前缀。

### 具体解法


#### **Golang**
```go
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res string
	sort.Strings(strs)
	for k := range strs[0] {
		flag := true
		for i := 0; i < len(strs)-1; i++ {
			if strs[i][k] != strs[i+1][k] {
				flag = false
				break
			}
		}
		if flag {
			res += strs[0][k : k+1]
		} else {
			break
		}
	}
	return res
}
```


# [187. 重复的DNA序列](https://leetcode-cn.com/problems/repeated-dna-sequences/description/)

### 题目描述

<p>所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：&ldquo;ACGAATTCCG&rdquo;。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。</p>

<p>编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 <code>s</code> 中出现次数超过一次。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>s = &quot;AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT&quot;
<strong>输出：</strong>[&quot;AAAAACCCCC&quot;, &quot;CCCCCAAAAA&quot;]</pre>

### 解题思路

1. hash table

### 具体解法


#### **Golang**
```go
func findRepeatedDnaSequences(s string) []string {
	sMap := make(map[string]bool)
	var res []string
	for i := 0; i < len(s)-9; i++ {
		if _, ok := sMap[s[i:i+10]]; ok {
			if sMap[s[i:i+10]] == true {
				res = append(res, s[i:i+10])
			}
			sMap[s[i:i+10]] = false
		} else {
			sMap[s[i:i+10]] = true
		}
	}
	return res
}

```



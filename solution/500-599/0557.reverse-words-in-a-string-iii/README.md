# [557. 反转字符串中的单词 III](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/)

### 题目描述

<p>给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre>
输入: &quot;Let&#39;s take LeetCode contest&quot;
输出: &quot;s&#39;teL ekat edoCteeL tsetnoc&quot;<strong><strong><strong>&nbsp;</strong></strong></strong>
</pre>

<p><strong><strong><strong><strong>注意：</strong></strong></strong></strong>在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。</p>

### 解题思路

1. 遍历字符串遇到空格就停止
2. 利用双指针 实现字符串的反转
3. 最后一个字符串是没有空格的，所以当时最后一个word的时候需要特殊处理一下结束的索引位置

### 具体解法


#### **Golang**
```go
func reverseWords(s string) string {
	b := []byte(s)
	l := 0
	for i, v := range s {
		if v == ' ' || i == len(s)-1 {
			r := i - 1
			if i == len(s)-1 {
				r = i
			}
			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}
			l = i + 1
		}
	}
	return string(b)
}
```



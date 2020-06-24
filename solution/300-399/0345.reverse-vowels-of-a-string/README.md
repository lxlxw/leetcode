# [345. 反转字符串中的元音字母](https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/)


### 题目描述

<p>编写一个函数，以字符串作为输入，反转该字符串中的元音字母。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入: </strong>&quot;hello&quot;
<strong>输出: </strong>&quot;holle&quot;
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入: </strong>&quot;leetcode&quot;
<strong>输出: </strong>&quot;leotcede&quot;</pre>

<p><strong>说明:</strong><br>
元音字母不包含字母&quot;y&quot;。</p>

### 解题思路

1. 双指针操作

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func reverseVowels(s string) string {
	vowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	p, q := 0, len(s)-1
	ret := []byte(s)
	for p < q {
		_, ok1 := vowel[unicode.ToLower(rune(ret[p]))]
		if !ok1 {
			p++
		}
		_, ok2 := vowel[unicode.ToLower(rune(ret[q]))]
		if !ok2 {
			q--
		}
		if ok1 && ok2 {
			ret[p], ret[q] = ret[q], ret[p]
			p++
			q--
		}
	}
	return string(ret)
}
```


<!-- tabs:end -->
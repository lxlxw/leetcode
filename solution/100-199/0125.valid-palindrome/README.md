# [125.验证回文串](https://leetcode-cn.com/problems/valid-palindrome)


### 题目描述

<div class="notranslate"><p>给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。</p>

<p><strong>说明：</strong>本题中，我们将空字符串定义为有效的回文串。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> "A man, a plan, a canal: Panama"
<strong>输出:</strong> true
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> "race a car"
<strong>输出:</strong> false
</pre>
</div>

### 解题思路

![](http://lc-photo.xwlin.com/125.png)

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9') {
			l++
		}
		for l < r && (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9') {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}
```


<!-- tabs:end -->
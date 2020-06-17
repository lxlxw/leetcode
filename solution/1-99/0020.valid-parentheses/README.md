# [20.有效的括号](https://leetcode-cn.com/problems/valid-parentheses)


### 题目描述

<div class="notranslate"><p>给定一个只包括 <code>'('</code>，<code>')'</code>，<code>'{'</code>，<code>'}'</code>，<code>'['</code>，<code>']'</code>&nbsp;的字符串，判断字符串是否有效。</p>

<p>有效字符串需满足：</p>

<ol>
	<li>左括号必须用相同类型的右括号闭合。</li>
	<li>左括号必须以正确的顺序闭合。</li>
</ol>

<p>注意空字符串可被认为是有效字符串。</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> "()"
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> "()[]{}"
<strong>输出:</strong> true
</pre>

<p><strong>示例&nbsp;3:</strong></p>

<pre><strong>输入:</strong> "(]"
<strong>输出:</strong> false
</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre><strong>输入:</strong> "([)]"
<strong>输出:</strong> false
</pre>

<p><strong>示例&nbsp;5:</strong></p>

<pre><strong>输入:</strong> "{[]}"
<strong>输出:</strong> true</pre>
</div>

### 解题思路
![](http://lc-photo.xwlin.com/20.gif)

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
import (
	"container/list"
	"strings"
)

func isValid(s string) bool {
	if s == "" {
		return true
	}
	s = strings.Replace(s, " ", "", -1)
	stack := list.New()
	trMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for v := range s {
		switch s[v : v+1] {
		case "(", "{", "[":
			stack.PushBack(s[v : v+1])
		case ")", "}", "]":
			if stack.Len() == 0 || stack.Back().Value != trMap[s[v:v+1]] {
				return false
			} else {
				stack.Remove(stack.Back())
			}
		}
	}
	return stack.Len() == 0
}
```
#### **Java**

#### **Python**

#### **PHP**

<!-- tabs:end -->
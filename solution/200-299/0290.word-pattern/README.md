# [290. 单词规律](https://leetcode-cn.com/problems/word-pattern/description/)

### 题目描述

<p>给定一种规律 <code>pattern</code>&nbsp;和一个字符串&nbsp;<code>str</code>&nbsp;，判断 <code>str</code> 是否遵循相同的规律。</p>

<p>这里的&nbsp;<strong>遵循&nbsp;</strong>指完全匹配，例如，&nbsp;<code>pattern</code>&nbsp;里的每个字母和字符串&nbsp;<code>str</code><strong>&nbsp;</strong>中的每个非空单词之间存在着双向连接的对应规律。</p>

<p><strong>示例1:</strong></p>

<pre><strong>输入:</strong> pattern = <code>&quot;abba&quot;</code>, str = <code>&quot;dog cat cat dog&quot;</code>
<strong>输出:</strong> true</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong>pattern = <code>&quot;abba&quot;</code>, str = <code>&quot;dog cat cat fish&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>示例 3:</strong></p>

<pre><strong>输入:</strong> pattern = <code>&quot;aaaa&quot;</code>, str = <code>&quot;dog cat cat dog&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>示例&nbsp;4:</strong></p>

<pre><strong>输入:</strong> pattern = <code>&quot;abba&quot;</code>, str = <code>&quot;dog dog dog dog&quot;</code>
<strong>输出:</strong> false</pre>

<p><strong>说明:</strong><br>
你可以假设&nbsp;<code>pattern</code>&nbsp;只包含小写字母，&nbsp;<code>str</code>&nbsp;包含了由单个空格分隔的小写字母。&nbsp; &nbsp;&nbsp;</p>

### 解题思路

两个哈希，一个存储pattern对str的关系，一个存储str对pattern的关系，即：
1. pMap以pattern的字符为key，str的单词为value；qMap则相反
2. 遍历str字符串数组，同步比较相同索引位置的pattern字符。
3. 如果遇到的pattern字符pMap中已经存在，取出哈希值，比较相同索引处的str单词，如果不同，返回false
4. 如果遇到的pattern字符pMap中不存在，则为新的字符，这时候以相同索引处的str单词为key，看看qMap中有没有值，
如果有，继续比较该str单词的pattern字符，如果不同，返回false
5. 经过上面的两个考验，还活了下来，说明目前的模式匹配没问题，那就分别设置pMap和qMap，继续循环就好了
6. 所有位置比较过了，都没问题的话，就返回true吧


### 具体解法


#### **Golang**
```go
func wordPattern(pattern string, str string) bool {

	sArr := strings.Split(str, " ")
	if len(sArr) != len(pattern) {
		return false
	}
	pMap := make(map[byte]string)
	qMap := make(map[string]byte)

	for i, v := range sArr {
		if _, ok := pMap[pattern[i]]; ok && pMap[pattern[i]] != v {
			return false
		} else {
			pMap[pattern[i]] = v
		}
		if _, ok := qMap[v]; ok && qMap[v] != pattern[i] {
			return false
		} else {
			qMap[v] = pattern[i]
		}
	}
	return true
}
```



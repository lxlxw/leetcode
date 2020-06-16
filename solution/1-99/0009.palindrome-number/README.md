# [9.回文数](https://leetcode-cn.com/problems/palindrome-number)


### 题目描述

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

**示例1:**

```json
输入: 121
输出: true
```

**示例2:**

```json
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

**示例3:**

```json
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
```

### 解题思路
![](http://lc-photo.xwlin.com/9.gif)

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	len := len(str)
	for i := 0; i < len; i++ {
		if str[i:i+1] != str[len-i-1:len-i] {
			return false
		}
	}
	return true
}
```
#### **Java**

#### **Python**

#### **PHP**

<!-- tabs:end -->
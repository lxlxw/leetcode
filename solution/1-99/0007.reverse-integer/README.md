# [7.整数反转](https://leetcode-cn.com/problems/reverse-integer)


### 题目描述

给出一个 `32` 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

**示例1:**

```json
输入: 123
输出: 321
```

**示例2:**

```json
输入: -123
输出: -321
```

**示例3:**

```json
输入: 120
输出: 21
```

**注意:**

假设我们的环境只能存储得下 `32` 位的有符号整数，则其数值范围为 `[−231,  231 − 1]`。请根据这个假设，如果反转后整数溢出那么就返回 `0`。

### 解题思路
![](https://pic.leetcode-cn.com/a161e3f0d41ad2866cb2bffe12084963b4e4c2e20086b71e377618d6b944fe4c-image.png)

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
import "math"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		if ret > math.MaxInt32 || ret < -math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	return ret
}
```
#### **Java**

#### **Python**

#### **PHP**

<!-- tabs:end -->
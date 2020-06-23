# [118.杨辉三角](https://leetcode-cn.com/problems/pascals-triangle/description/)


### 题目描述

<div class="notranslate"><p>给定一个非负整数&nbsp;<em>numRows，</em>生成杨辉三角的前&nbsp;<em>numRows&nbsp;</em>行。</p>

<p><img src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif" alt=""></p>

<p><small>在杨辉三角中，每个数是它左上方和右上方的数的和。</small></p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> 5
<strong>输出:</strong>
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]</pre>
</div>

### 解题思路

1. 递归实现

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	var res [][]int
	res = append(res, []int{1})
	res = append(res, []int{1, 1})
	recursive(2, numRows, &res)
	return res
}

func recursive(current, numRows int, res *[][]int) {
	row := make([]int, current+1)
	row[0], row[current] = 1, 1
	for i := 1; i < current; i++ {
		row[i] = (*res)[current-1][i-1] + (*res)[current-1][i]
	}
	*res = append(*res, row)
	if current+1 == numRows {
		return
	} else {
		recursive(current+1, numRows, res)
	}
}
```


<!-- tabs:end -->
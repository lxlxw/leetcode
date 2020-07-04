# [48. 旋转图像](https://leetcode-cn.com/problems/rotate-image)

### 题目描述

<p>给定一个 <em>n&nbsp;</em>&times;&nbsp;<em>n</em> 的二维矩阵表示一个图像。</p>

<p>将图像顺时针旋转 90 度。</p>

<p><strong>说明：</strong></p>

<p>你必须在<strong><a href="https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95" target="_blank">原地</a></strong>旋转图像，这意味着你需要直接修改输入的二维矩阵。<strong>请不要</strong>使用另一个矩阵来旋转图像。</p>

<p><strong>示例 1:</strong></p>

<pre>给定 <strong>matrix</strong> = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

<strong>原地</strong>旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
</pre>

<p><strong>示例 2:</strong></p>

<pre>给定 <strong>matrix</strong> =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

<strong>原地</strong>旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
</pre>



### 解题思路

![](http://lc-photo.xwlin.com/48-1.png)
![](http://lc-photo.xwlin.com/48-2.png)
![](http://lc-photo.xwlin.com/48-3.png)
![](http://lc-photo.xwlin.com/48-4.png)
![](http://lc-photo.xwlin.com/48-5.png)

### 具体解法


#### **Golang**
```go
func rotate(matrix [][]int) {
	size := len(matrix)
	subsize := size - 1
	for i := 0; i < size/2; i++ {
		for j := i; j < subsize-i; j++ {
			matrix[i][j], matrix[j][subsize-i] = matrix[j][subsize-i], matrix[i][j]
			matrix[i][j], matrix[subsize-i][subsize-j] = matrix[subsize-i][subsize-j], matrix[i][j]
			matrix[i][j], matrix[subsize-j][i] = matrix[subsize-j][i], matrix[i][j]
		}
	}
}
```


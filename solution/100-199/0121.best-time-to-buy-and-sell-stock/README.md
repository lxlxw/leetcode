# [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/)

### 题目描述

<p>给定一个数组，它的第&nbsp;<em>i</em> 个元素是一支给定股票第 <em>i</em> 天的价格。</p>

<p>如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。</p>

<p>注意：你不能在买入股票前卖出股票。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [7,1,5,3,6,4]
<strong>输出:</strong> 5
<strong>解释: </strong>在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> [7,6,4,3,1]
<strong>输出:</strong> 0
<strong>解释: </strong>在这种情况下, 没有交易完成, 所以最大利润为 0。
</pre>

### 解题思路

 1. 设定`buy`为 prices[0], `sale`, `max`为0
 2. 遍历剩余的数组元素
 3. 判断当前数组元素`v`减去`buy`大于0时，表示利润大于0，此时表示可以出售，赋值给`sale`
 4. 判断本次出售的`sale`是否大于上一次出售的`max`，如果大于，则更新`max`
 5. 判断当前数组元素`v`是否小于`buy`，如果小于，则表示更新`buy`

### 具体解法


#### **Golang**
```go
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	sale, max := 0, 0
	for _, v := range prices[1:] {
		if v-buy > 0 {
			sale = v - buy
		}
		if sale > max {
			max = sale
		}
		if v < buy {
			buy = v
		}
	}
	return max
}
```



# [299. Nim 游戏](https://leetcode-cn.com/problems/nim-game/description/)

### 题目描述

<p>你在和朋友一起玩 <a href="https://baike.baidu.com/item/%E7%8C%9C%E6%95%B0%E5%AD%97/83200?fromtitle=Bulls+and+Cows&amp;fromid=12003488&amp;fr=aladdin" target="_blank">猜数字（Bulls and Cows）</a>游戏，该游戏规则如下：</p>

<ol>
	<li>你写出一个秘密数字，并请朋友猜这个数字是多少。</li>
	<li>朋友每猜测一次，你就会给他一个提示，告诉他的猜测数字中有多少位属于数字和确切位置都猜对了（称为&ldquo;Bulls&rdquo;, 公牛），有多少位属于数字猜对了但是位置不对（称为&ldquo;Cows&rdquo;, 奶牛）。</li>
	<li>朋友根据提示继续猜，直到猜出秘密数字。</li>
</ol>

<p>请写出一个根据秘密数字和朋友的猜测数返回提示的函数，返回字符串的格式为 <code>xAyB</code> ，<code>x</code> 和 <code>y</code> 都是数字，<code>A</code> 表示公牛，用&nbsp;<code>B</code>&nbsp;表示奶牛。</p>

<ul>
	<li><code>xA</code> 表示有 <code>x</code> 位数字出现在秘密数字中，且位置都与秘密数字一致。</li>
	<li><code>yB</code> 表示有 <code>y</code> 位数字出现在秘密数字中，但位置与秘密数字不一致。</li>
</ul>

<p>请注意秘密数字和朋友的猜测数都可能含有重复数字，每位数字只能统计一次。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> secret = &quot;1807&quot;, guess = &quot;7810&quot;
<strong>输出:</strong> &quot;1A3B&quot;
<strong>解释:</strong> <code>1</code>&nbsp;公牛和&nbsp;<code>3</code>&nbsp;奶牛。公牛是 <code>8</code>，奶牛是 <code>0</code>, <code>1</code>&nbsp;和 <code>7</code>。</pre>

<p><strong>示例 2:</strong></p>

<pre><strong>输入:</strong> secret = &quot;1123&quot;, guess = &quot;0111&quot;
<strong>输出:</strong> &quot;1A1B&quot;
<strong>解释: </strong>朋友猜测数中的第一个 <code>1</code>&nbsp;是公牛，第二个或第三个 <code>1</code>&nbsp;可被视为奶牛。</pre>

<p>&nbsp;</p>

<p><strong>说明: </strong>你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。</p>

### 解题思路



### 具体解法


#### **Golang**
```go
func getHint(secret string, guess string) string {
	var countS, countG [10]int
	bulls, cows := 0, 0
	for i := range secret {
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')
		if ng == ns {
			bulls++
			continue
		}

		if countG[ns] > 0 {
			cows++
			countG[ns]--
		} else {
			countS[ns]++
		}

		if countS[ng] > 0 {
			cows++
			countS[ng]--
		} else {
			countG[ng]++
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
```



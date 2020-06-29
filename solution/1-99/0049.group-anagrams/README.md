# [49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams)

### 题目描述

<p>给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> <code>[&quot;eat&quot;, &quot;tea&quot;, &quot;tan&quot;, &quot;ate&quot;, &quot;nat&quot;, &quot;bat&quot;]</code>,
<strong>输出:</strong>
[
  [&quot;ate&quot;,&quot;eat&quot;,&quot;tea&quot;],
  [&quot;nat&quot;,&quot;tan&quot;],
  [&quot;bat&quot;]
]</pre>

<p><strong>说明：</strong></p>

<ul>
	<li>所有输入均为小写字母。</li>
	<li>不考虑答案输出的顺序。</li>
</ul>



### 解题思路

1. 对字符串排序
2. 利用hash table判断是否存在,如果已存在,则将该字符串分配到已存在的返回数组中

### 具体解法

#### **Golang**
```go
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string]int)
	var res [][]string
	mapIndex := 0
	for _, str := range strs {
		strSort := sortString(str)

		if _, ok := strMap[strSort]; !ok {
			res = append(res, []string{str})
			strMap[strSort] = mapIndex
			mapIndex++
		} else {
			res[strMap[strSort]] = append(res[strMap[strSort]], str)
		}
	}
	return res
}

func sortString(str string) string {
	strR := []rune(str)
	strArr := []string{}
	for i := 0; i < len(strR); i++ {
		strArr = append(strArr, string(strR[i]))
	}
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}
```


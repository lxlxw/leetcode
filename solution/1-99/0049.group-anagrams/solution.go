package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
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

// @lc code=end

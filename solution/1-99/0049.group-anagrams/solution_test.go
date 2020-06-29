package leetcode

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var ret [][]string
	var strs []string

	strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ret = [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	for k, str := range groupAnagrams(strs) {
		for i, v := range str {
			if ret[k][i] != v {
				t.Fatalf("case fails %v\n", ret)
			}
		}
	}
}

package leetcode

import (
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	romanManyMap := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	var ret int
	for k, v := range romanManyMap {
		if strings.Contains(s, k) {
			s = strings.Replace(s, k, strconv.Itoa(v), -1)
			ret += v
		}
	}
	for k := range s {
		if romanMap[s[k:k+1]] != 0 {
			ret += romanMap[s[k:k+1]]
		}
	}
	return ret
}

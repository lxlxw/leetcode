package leetcode

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {

	var s string
	var ret int
	s = "III"
	ret = 3
	fmt.Printf("x = %v ret = %v\n", s, romanToInt(s))
	if romanToInt(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = "IV"
	ret = 4
	fmt.Printf("x = %v ret = %v\n", s, romanToInt(s))
	if romanToInt(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = "IX"
	ret = 9
	fmt.Printf("x = %v ret = %v\n", s, romanToInt(s))
	if romanToInt(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = "LVIII"
	ret = 58
	fmt.Printf("x = %v ret = %v\n", s, romanToInt(s))
	if romanToInt(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = "MCMXCIV"
	ret = 1994
	fmt.Printf("x = %v ret = %v\n", s, romanToInt(s))
	if romanToInt(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}

package leetcode

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	var ret bool
	var num int
	num = 19
	ret = true
	if ret != isHappy(num) {
		t.Fatalf("case fails %v\n", ret)
	}

	num = 2
	ret = false
	if ret != isHappy(num) {
		t.Fatalf("case fails %v\n", ret)
	}
}
